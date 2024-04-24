package avsregistry

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	contractOperatorStateRetriever "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/OperatorStateRetriever"
	opstateretriever "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/StakeRegistry"
	exoabi "github.com/ExocoreNetwork/exocore-sdk/types/abi"
)

type AvsRegistryReader interface {
	GetCheckSignaturesIndices(
		opts *bind.CallOpts,
		referenceBlockNumber uint32,
		quorumNumbers []byte,
		nonSignerOperatorIds [][32]byte,
	) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorId(opts *bind.CallOpts, operatorAddress gethcommon.Address) ([32]byte, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId types.OperatorId) (gethcommon.Address, error)

	IsOperatorRegistered(opts *bind.CallOpts, operatorAddress gethcommon.Address) (bool, error)

	QueryExistingRegisteredOperatorPubKeys(
		ctx context.Context,
		startBlock *big.Int,
		stopBlock *big.Int,
	) ([]types.OperatorAddr, []types.OperatorPubkeys, error)
}

type AvsRegistryChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      gethcommon.Address
	registryCoordinatorAddr gethcommon.Address
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.EthClient
}

// forces AvsReader to implement the clients.ReaderInterface interface
var _ AvsRegistryReader = (*AvsRegistryChainReader)(nil)

func NewAvsRegistryChainReader(
	registryCoordinatorAddr gethcommon.Address,
	blsApkRegistryAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	logger logging.Logger,
	ethClient eth.EthClient,
) *AvsRegistryChainReader {
	return &AvsRegistryChainReader{
		blsApkRegistryAddr:      blsApkRegistryAddr,
		registryCoordinatorAddr: registryCoordinatorAddr,
		registryCoordinator:     registryCoordinator,
		operatorStateRetriever:  operatorStateRetriever,
		stakeRegistry:           stakeRegistry,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

func BuildAvsRegistryChainReader(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryChainReader, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, err
	}
	blsApkRegistryAddr, err := contractRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := contractRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	contractOperatorStateRetriever, err := contractOperatorStateRetriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
	}
	return NewAvsRegistryChainReader(
		registryCoordinatorAddr,
		blsApkRegistryAddr,
		contractRegistryCoordinator,
		contractOperatorStateRetriever,
		contractStakeRegistry,
		logger,
		ethClient,
	), nil
}

func (r *AvsRegistryChainReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers []byte,
	nonSignerOperatorIds [][32]byte,
) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	checkSignatureIndices, err := r.operatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		r.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers,
		nonSignerOperatorIds,
	)
	if err != nil {
		r.logger.Error("Failed to get check signatures indices", "err", err)
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, err
	}
	return checkSignatureIndices, nil
}

func (r *AvsRegistryChainReader) GetOperatorId(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) ([32]byte, error) {
	operatorId, err := r.registryCoordinator.GetOperatorId(
		opts,
		operatorAddress,
	)
	if err != nil {
		r.logger.Error("Failed to get operator id", "err", err)
		return [32]byte{}, err
	}
	return operatorId, nil
}

func (r *AvsRegistryChainReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (gethcommon.Address, error) {
	operatorAddress, err := r.registryCoordinator.GetOperatorFromId(
		opts,
		operatorId,
	)
	if err != nil {
		r.logger.Error("Failed to get operator address", "err", err)
		return gethcommon.Address{}, err
	}
	return operatorAddress, nil
}

func (r *AvsRegistryChainReader) IsOperatorRegistered(
	opts *bind.CallOpts,
	operatorAddress gethcommon.Address,
) (bool, error) {
	operatorStatus, err := r.registryCoordinator.GetOperatorStatus(opts, operatorAddress)
	if err != nil {
		r.logger.Error("Cannot get operator status", "err", err)
		return false, err
	}

	// 0 = NEVER_REGISTERED, 1 = REGISTERED, 2 = DEREGISTERED
	registeredWithAvs := operatorStatus == 1
	return registeredWithAvs, nil
}

func (r *AvsRegistryChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {

	blsApkRegistryAbi, err := abi.JSON(bytes.NewReader(exoabi.BLSApkRegistryAbi))
	if err != nil {
		r.logger.Error("Error getting Abi", "err", err)
		return nil, nil, err
	}

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   stopBlock,
		Addresses: []gethcommon.Address{
			r.blsApkRegistryAddr,
		},
		Topics: [][]gethcommon.Hash{{blsApkRegistryAbi.Events["NewPubkeyRegistration"].ID}},
	}

	logs, err := r.ethClient.FilterLogs(ctx, query)
	if err != nil {
		r.logger.Error("Error filtering logs", "err", err)
		return nil, nil, err
	}
	r.logger.Info("avsRegistryChainReader.QueryExistingRegisteredOperatorPubKeys", "transactionLogs", logs)

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)

	for _, vLog := range logs {

		// get the operator address
		operatorAddr := gethcommon.HexToAddress(vLog.Topics[1].Hex())
		operatorAddresses = append(operatorAddresses, operatorAddr)

		event, err := blsApkRegistryAbi.Unpack("NewPubkeyRegistration", vLog.Data)
		if err != nil {
			r.logger.Error("Error unpacking event data", "err", err)
			return nil, nil, err
		}

		G1Pubkey := event[0].(struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		})

		G2Pubkey := event[1].(struct {
			X [2]*big.Int "json:\"X\""
			Y [2]*big.Int "json:\"Y\""
		})

		operatorPubkey := types.OperatorPubkeys{
			G1Pubkey: bls.NewG1Point(
				G1Pubkey.X,
				G1Pubkey.Y,
			),
			G2Pubkey: bls.NewG2Point(
				G2Pubkey.X,
				G2Pubkey.Y,
			),
		}

		operatorPubkeys = append(operatorPubkeys, operatorPubkey)

	}

	return operatorAddresses, operatorPubkeys, nil
}
