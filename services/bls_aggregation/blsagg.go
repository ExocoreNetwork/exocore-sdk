package blsagg

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/services/avsregistry"
	"github.com/ExocoreNetwork/exocore-sdk/types"
	"github.com/ExocoreNetwork/exocore-sdk/utils"
)

var (
	// TODO: refactor these errors to use a custom struct with taskIndex field instead of wrapping taskIndex in the error string directly.
	//       see https://go.dev/blog/go1.13-errors
	TaskInitializationErrorFn = func(err error, taskIndex types.TaskIndex) error {
		return fmt.Errorf("Failed to initialize task %d: %w", taskIndex, err)
	}
	TaskAlreadyInitializedErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d already initialized", taskIndex)
	}
	TaskExpiredErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d expired", taskIndex)
	}
	TaskNotFoundErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	OperatorNotPartOfTaskQuorumErrorFn = func(operatorId types.OperatorId, taskIndex types.TaskIndex) error {
		return fmt.Errorf("operator %x not part of task %d's quorum", operatorId, taskIndex)
	}
	HashFunctionError = func(err error) error {
		return fmt.Errorf("Failed to hash task response: %w", err)
	}
	SignatureVerificationError = func(err error) error {
		return fmt.Errorf("Failed to verify signature: %w", err)
	}
	IncorrectSignatureError = errors.New("Signature verification failed. Incorrect Signature.")
)

// BlsAggregationServiceResponse is the response from the bls aggregation service
type BlsAggregationServiceResponse struct {
	Err                error                    // if Err is not nil, the other fields are not valid
	TaskIndex          types.TaskIndex          // unique identifier of the task
	TaskResponse       types.TaskResponse       // the task response that was signed
	TaskResponseDigest types.TaskResponseDigest // digest of the task response that was signed
	// The below 8 fields are the data needed to build the IBLSSignatureChecker.NonSignerStakesAndSignature struct
	// users of this service will need to build the struct themselves by converting the bls points
	// into the BN254.G1/G2Point structs that the IBLSSignatureChecker expects
	// given that those are different for each AVS service manager that individually inherits BLSSignatureChecker
	NonSignersPubkeysG1 []*bls.G1Point
	QuorumApksG1        []*bls.G1Point
	SignersApkG2        *bls.G2Point
	SignersAggSigG1     *bls.Signature
}

// aggregatedOperators is meant to be used as a value in a map
// map[taskResponseDigest]aggregatedOperators
type aggregatedOperators struct {
	// aggregate g2 pubkey of all operatos who signed on this taskResponseDigest
	signersApkG2 *bls.G2Point
	// aggregate signature of all operators who signed on this taskResponseDigest
	signersAggSigG1 *bls.Signature
	// aggregate stake of all operators who signed on this header for each quorum
	signersTotalStakePerQuorum map[types.QuorumNum]*big.Int
	// set of OperatorId of operators who signed on this header
	signersOperatorIdsSet map[types.OperatorId]bool
}

// BlsAggregationService is the interface provided to avs aggregator code for doing bls aggregation
// Currently its only implementation is the BlsAggregatorService, so see the comment there for more details
type BlsAggregationService interface {
	// InitializeNewTask should be called whenever a new task is created. ProcessNewSignature will return an error
	// if the task it is trying to process has not been initialized yet.
	// quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered complete, which happens
	// when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers whose stake
	// in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in that quorum
	InitializeNewTask(
		taskIndex types.TaskIndex,
		taskCreatedBlock uint32,
		quorumNumbers types.QuorumNums,
		quorumThresholdPercentages types.QuorumThresholdPercentages,
		timeToExpiry time.Duration,
	) error

	// ProcessNewSignature processes a new signature over a taskResponseDigest for a particular taskIndex by a particular operator
	// It verifies that the signature is correct and returns an error if it is not, and then aggregates the signature and stake of
	// the operator with all other signatures for the same taskIndex and taskResponseDigest pair.
	// Note: This function currently only verifies signatures over the taskResponseDigest directly, so avs code needs to verify that the digest
	// passed to ProcessNewSignature is indeed the digest of a valid taskResponse (that is, BlsAggregationService does not verify semantic integrity of the taskResponses)
	ProcessNewSignature(
		ctx context.Context,
		taskIndex types.TaskIndex,
		taskResponse types.TaskResponse,
		blsSignature *bls.Signature,
		operatorId types.OperatorId,
	) error

	// GetResponseChannel returns the single channel that meant to be used as the response channel
	// Any task that is completed (see the completion criterion in the comment above InitializeNewTask)
	// will be sent on this channel along with all the necessary information to call BLSSignatureChecker onchain
	GetResponseChannel() <-chan BlsAggregationServiceResponse
}

// BlsAggregatorService is a service that performs BLS signature aggregation for an AVS' tasks
// Assumptions:
//  1. BlsAggregatorService only verifies digest signatures, so avs code needs to verify that the digest
//     passed to ProcessNewSignature is indeed the digest of a valid taskResponse
//     (see the comment above checkSignature for more details)
//  2. BlsAggregatorService is VERY generic and makes very few assumptions about the tasks structure or
//     the time at which operators will send their signatures. It is mostly suitable for offchain computation
//     oracle (a la truebit) type of AVS, where tasks are sent onchain by users sporadically, and where
//     new tasks can start even before the previous ones have finished aggregation.
//     AVSs like eigenDA that have a much more controlled task submission schedule and where new tasks are
//     only submitted after the previous one's response has been aggregated and responded onchain, could have
//     a much simpler AggregationService without all the complicated parallel goroutines.
type BlsAggregatorService struct {
	// aggregatedResponsesC is the channel which all goroutines share to send their responses back to the
	// main thread after they are done aggregating (either they reached the threshold, or timeout expired)
	aggregatedResponsesC chan BlsAggregationServiceResponse
	// signedTaskRespsCs are the channels to send the signed task responses to the goroutines processing them
	// each new task is assigned a new goroutine and a new channel
	signedTaskRespsCs map[types.TaskIndex]chan types.SignedTaskResponseDigest
	// we add chans to taskChans from the main thread (InitializeNewTask) when we create new tasks,
	// we read them in ProcessNewSignature from the main thread when we receive new signed tasks,
	// and remove them from its respective goroutine when the task is completed or reached timeout
	// we thus need a mutex to protect taskChans
	taskChansMutex     sync.RWMutex
	avsRegistryService avsregistry.FakeAvsRegistryService
	logger             logging.Logger

	hashFunction types.TaskResponseHashFunction
}

var _ BlsAggregationService = (*BlsAggregatorService)(nil)

func NewBlsAggregatorService(avsRegistryService avsregistry.FakeAvsRegistryService, hashFunction types.TaskResponseHashFunction, logger logging.Logger) *BlsAggregatorService {
	return &BlsAggregatorService{
		aggregatedResponsesC: make(chan BlsAggregationServiceResponse),
		signedTaskRespsCs:    make(map[types.TaskIndex]chan types.SignedTaskResponseDigest),
		taskChansMutex:       sync.RWMutex{},
		avsRegistryService:   avsRegistryService,
		logger:               logger,
		hashFunction:         hashFunction,
	}
}

func (a *BlsAggregatorService) GetResponseChannel() <-chan BlsAggregationServiceResponse {
	return a.aggregatedResponsesC
}

// InitializeNewTask creates a new task goroutine meant to process new signed task responses for that task
// (that are sent via ProcessNewSignature) and adds a channel to a.taskChans to send the signed task responses to it
// quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered complete, which happens
// when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers whose stake
// in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in that quorum
func (a *BlsAggregatorService) InitializeNewTask(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers types.QuorumNums,
	quorumThresholdPercentages types.QuorumThresholdPercentages,
	timeToExpiry time.Duration,
) error {
	a.logger.Debug("AggregatorService initializing new task", "taskIndex", taskIndex, "taskCreatedBlock", taskCreatedBlock, "quorumNumbers", quorumNumbers, "quorumThresholdPercentages", quorumThresholdPercentages, "timeToExpiry", timeToExpiry)

	a.taskChansMutex.Lock()
	defer a.taskChansMutex.Unlock()
	if _, taskExists := a.signedTaskRespsCs[taskIndex]; taskExists {
		return TaskAlreadyInitializedErrorFn(taskIndex)
	}
	signedTaskRespsC := make(chan types.SignedTaskResponseDigest)
	a.signedTaskRespsCs[taskIndex] = signedTaskRespsC

	go a.singleTaskAggregatorGoroutineFunc(taskIndex, taskCreatedBlock, quorumNumbers, quorumThresholdPercentages, timeToExpiry, signedTaskRespsC)
	return nil
}

func (a *BlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskResponse types.TaskResponse,
	blsSignature *bls.Signature,
	operatorId types.OperatorId,
) error {
	a.taskChansMutex.Lock()
	taskC, taskInitialized := a.signedTaskRespsCs[taskIndex]
	a.taskChansMutex.Unlock()
	if !taskInitialized {
		return TaskNotFoundErrorFn(taskIndex)
	}

	signatureVerificationErrorC := make(chan error)
	// send the task to the goroutine processing this task
	// and return the error (if any) returned by the signature verification routine

	select {
	// we need to send this as part of select because if the goroutine is processing another SignedTaskResponseDigest
	// and cannot receive this one, we want the context to be able to cancel the request
	case taskC <- types.SignedTaskResponseDigest{
		TaskResponse:                taskResponse,
		BlsSignature:                blsSignature,
		OperatorId:                  operatorId,
		SignatureVerificationErrorC: signatureVerificationErrorC,
	}:
		// note that we need to wait synchronously here for this response because we want to
		// send back an informative error message to the operator who sent his signature to the aggregator
		return <-signatureVerificationErrorC
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *BlsAggregatorService) singleTaskAggregatorGoroutineFunc(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers types.QuorumNums,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	signedTaskRespsC <-chan types.SignedTaskResponseDigest,
) {
	defer a.closeTaskGoroutine(taskIndex)
	quorumThresholdPercentagesMap := make(map[types.QuorumNum]types.QuorumThresholdPercentage)
	for i, quorumNumber := range quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = quorumThresholdPercentages[i]
	}
	operatorsAvsStateDict, err := a.avsRegistryService.GetOperatorsAvsStateAtBlock(context.Background(), quorumNumbers, taskCreatedBlock)
	if err != nil {
		a.logger.Error("Task goroutine failed to get operators state from avs registry", "taskIndex", taskIndex, "err", err)
		a.aggregatedResponsesC <- BlsAggregationServiceResponse{
			Err:       TaskInitializationErrorFn(fmt.Errorf("AggregatorService failed to get operators state from avs registry at blockNum %d: %w", taskCreatedBlock, err), taskIndex),
			TaskIndex: taskIndex,
		}
		return
	}
	quorumsAvsStakeDict, err := a.avsRegistryService.GetQuorumsAvsStateAtBlock(context.Background(), quorumNumbers, taskCreatedBlock)
	if err != nil {
		a.logger.Error("Task goroutine failed to get quorums state from avs registry", "taskIndex", taskIndex, "err", err)
		a.aggregatedResponsesC <- BlsAggregationServiceResponse{
			Err:       TaskInitializationErrorFn(fmt.Errorf("Aggregator failed to get quorums state from avs registry: %w", err), taskIndex),
			TaskIndex: taskIndex,
		}
		return
	}
	totalStakePerQuorum := make(map[types.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
	}
	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	// TODO(samlaf): instead of taking a TTE, we should take a block as input
	// and monitor the chain and only close the task goroutine when that block is reached
	taskExpiredTimer := time.NewTimer(timeToExpiry)

	aggregatedOperatorsDict := map[types.TaskResponseDigest]aggregatedOperators{}
	for {
		select {
		case signedTaskResponseDigest := <-signedTaskRespsC:
			a.logger.Debug("Task goroutine received new signed task response digest", "taskIndex", taskIndex, "signedTaskResponseDigest", signedTaskResponseDigest)

			err := a.verifySignature(taskIndex, signedTaskResponseDigest, operatorsAvsStateDict)
			signedTaskResponseDigest.SignatureVerificationErrorC <- err
			if err != nil {
				continue
			}

			// compute the taskResponseDigest using the hash function
			taskResponseDigest, err := a.hashFunction(signedTaskResponseDigest.TaskResponse)
			if err != nil {
				// this error should never happen, because we've already hashed the taskResponse in verifySignature,
				// but keeping here in case the verifySignature implementation ever changes or some catastrophic bug happens..
				continue
			}
			// after verifying signature we aggregate its sig and pubkey, and update the signed stake amount
			digestAggregatedOperators, ok := aggregatedOperatorsDict[taskResponseDigest]
			if !ok {
				// first operator to sign on this digest
				digestAggregatedOperators = aggregatedOperators{
					// we've already verified that the operator is part of the task's quorum, so we don't need checks here
					signersApkG2:               bls.NewZeroG2Point().Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey),
					signersAggSigG1:            signedTaskResponseDigest.BlsSignature,
					signersOperatorIdsSet:      map[types.OperatorId]bool{signedTaskResponseDigest.OperatorId: true},
					signersTotalStakePerQuorum: cloneStakePerQuorumMap(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum),
				}
			} else {
				digestAggregatedOperators.signersAggSigG1.Add(signedTaskResponseDigest.BlsSignature)
				digestAggregatedOperators.signersApkG2.Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey)
				digestAggregatedOperators.signersOperatorIdsSet[signedTaskResponseDigest.OperatorId] = true
				for quorumNum, stake := range operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum {
					if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
						// if we haven't seen this quorum before, initialize its signed stake to 0
						// possible if previous operators who sent us signatures were not part of this quorum
						digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
					}
					digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
				}
			}
			// update the aggregatedOperatorsDict. Note that we need to assign the whole struct value at once,
			// because of https://github.com/golang/go/issues/3117
			aggregatedOperatorsDict[taskResponseDigest] = digestAggregatedOperators

			if checkIfStakeThresholdsMet(a.logger, digestAggregatedOperators.signersTotalStakePerQuorum, totalStakePerQuorum, quorumThresholdPercentagesMap) {
				nonSignersOperatorIds := []types.OperatorId{}
				for operatorId := range operatorsAvsStateDict {
					if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
						nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
					}
				}

				// the contract requires a sorted nonSignersOperatorIds
				sort.SliceStable(nonSignersOperatorIds, func(i, j int) bool {
					iOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[i][:])
					jOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[j][:])
					return iOprInt.Cmp(jOprInt) == -1
				})

				nonSignersG1Pubkeys := []*bls.G1Point{}
				for _, operatorId := range nonSignersOperatorIds {
					operator := operatorsAvsStateDict[operatorId]
					nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.OperatorInfo.Pubkeys.G1Pubkey)
				}

				if err != nil {
					a.aggregatedResponsesC <- BlsAggregationServiceResponse{
						Err:       utils.WrapError(errors.New("Failed to get check signatures indices"), err),
						TaskIndex: taskIndex,
					}
					return
				}

				blsAggregationServiceResponse := BlsAggregationServiceResponse{
					Err:                 nil,
					TaskIndex:           taskIndex,
					TaskResponse:        signedTaskResponseDigest.TaskResponse,
					TaskResponseDigest:  taskResponseDigest,
					NonSignersPubkeysG1: nonSignersG1Pubkeys,
					QuorumApksG1:        quorumApksG1,
					SignersApkG2:        digestAggregatedOperators.signersApkG2,
					SignersAggSigG1:     digestAggregatedOperators.signersAggSigG1,
				}
				a.aggregatedResponsesC <- blsAggregationServiceResponse
				taskExpiredTimer.Stop()
				return
			}
		case <-taskExpiredTimer.C:
			a.aggregatedResponsesC <- BlsAggregationServiceResponse{
				Err:       TaskExpiredErrorFn(taskIndex),
				TaskIndex: taskIndex,
			}
			return
		}
	}

}

// closeTaskGoroutine is run when the goroutine processing taskIndex's task responses ends (for whatever reason)
// it deletes the response channel for taskIndex from a.taskChans
// so that the main thread knows that this task goroutine is no longer running
// and doesn't try to send new signatures to it
func (a *BlsAggregatorService) closeTaskGoroutine(taskIndex types.TaskIndex) {
	a.taskChansMutex.Lock()
	delete(a.signedTaskRespsCs, taskIndex)
	a.taskChansMutex.Unlock()
}

// verifySignature verifies that a signature is valid against the operator pubkey stored in the
// operatorsAvsStateDict for that particular task
// TODO(samlaf): right now we are only checking that the *digest* is signed correctly!!
// we could be sent a signature of any kind of garbage and we would happily aggregate it
// this forces the avs code to verify that the digest is indeed the digest of a valid taskResponse
// we could take taskResponse as an interface{} and have avs code pass us a taskResponseHashFunction
// that we could use to hash and verify the taskResponse itself
func (a *BlsAggregatorService) verifySignature(
	taskIndex types.TaskIndex,
	signedTaskResponseDigest types.SignedTaskResponseDigest,
	operatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId]
	if !ok {
		a.logger.Warnf("Operator %#v not found. Skipping message.", signedTaskResponseDigest.OperatorId)
		return OperatorNotPartOfTaskQuorumErrorFn(signedTaskResponseDigest.OperatorId, taskIndex)
	}

	taskResponseDigest, err := a.hashFunction(signedTaskResponseDigest.TaskResponse)
	if err != nil {
		a.logger.Error("Failed to hash task response, skipping.", "taskIndex", taskIndex, "signedTaskResponseDigest", signedTaskResponseDigest, "err", err)
		return HashFunctionError(err)
	}

	// verify that the msg actually came from the correct operator
	operatorG2Pubkey := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		a.logger.Error("Operator G2 pubkey not found", "operatorId", signedTaskResponseDigest.OperatorId, "taskId", taskIndex)
		return fmt.Errorf("taskId %d: Operator G2 pubkey not found (operatorId: %x)", taskIndex, signedTaskResponseDigest.OperatorId)
	}
	a.logger.Debug("Verifying signed task response digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"taskResponseDigest", taskResponseDigest,
		"blsSignature", signedTaskResponseDigest.BlsSignature,
	)

	// if the operator signs a digest that is not the digest of the TaskResponse submitted in ProcessNewTask
	// then the signature will not be verified
	signatureVerified, err := signedTaskResponseDigest.BlsSignature.Verify(operatorG2Pubkey, taskResponseDigest)
	if err != nil {
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		return IncorrectSignatureError
	}
	return nil
}

// checkIfStakeThresholdsMet checks at least quorumThresholdPercentage of stake
// has signed for each quorum.
func checkIfStakeThresholdsMet(
	logger logging.Logger,
	signedStakePerQuorum map[types.QuorumNum]*big.Int,
	totalStakePerQuorum map[types.QuorumNum]*big.Int,
	quorumThresholdPercentagesMap map[types.QuorumNum]types.QuorumThresholdPercentage,
) bool {
	for quorumNum, quorumThresholdPercentage := range quorumThresholdPercentagesMap {
		signedStakeByQuorum, ok := signedStakePerQuorum[quorumNum]
		if !ok {
			// signedStakePerQuorum not contain the quorum,
			// this case means signedStakePerQuorum has not signed for each quorum.
			// even the total stake for this quorum is zero.
			return false
		}

		totalStakeByQuorum, ok := totalStakePerQuorum[quorumNum]
		if !ok {
			// Note this case should not happen.
			// The `totalStakePerQuorum` is got from the contract, so if we not found the
			// totalStakeByQuorum, that means the code have a bug.
			logger.Errorf("TotalStake not found for quorum %d.", quorumNum)
			return false
		}

		// we check that signedStake >= totalStake * quorumThresholdPercentage / 100
		// to be exact (and do like the contracts), we actually check that
		// signedStake * 100 >= totalStake * quorumThresholdPercentage
		signedStake := big.NewInt(0).Mul(signedStakeByQuorum, big.NewInt(100))
		thresholdStake := big.NewInt(0).Mul(totalStakeByQuorum, big.NewInt(int64(quorumThresholdPercentage)))
		if signedStake.Cmp(thresholdStake) < 0 {
			return false
		}
	}
	return true
}

func cloneStakePerQuorumMap(stakes map[types.QuorumNum]types.StakeAmount) map[types.QuorumNum]types.StakeAmount {
	out := make(map[types.QuorumNum]types.StakeAmount, len(stakes))
	for k, v := range stakes {
		out[k] = new(big.Int).Set(v)
	}
	return out
}
