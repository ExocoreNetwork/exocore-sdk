package avsregistry

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exocontracts"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	blsapkregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/StakeRegistry"
)

type AvsRegistryWriter interface {
	RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		ctx context.Context,
		operatorEcdsaPrivateKey *ecdsa.PrivateKey,
		operatorToAvsRegistrationSigSalt [32]byte,
		operatorToAvsRegistrationSigExpiry *big.Int,
		blsKeyPair *bls.KeyPair,
		quorumNumbers []byte,
		socket string,
	) (*gethtypes.Receipt, error)

	// UpdateStakesOfEntireOperatorSetForQuorums is used by avs teams running https://github.com/Layr-Labs/avs-sync
	// to updates the stake of their entire operator set.
	// Because of high gas costs of this operation, it typically needs to be called for every quorum, or perhaps for a
	// small grouping of quorums
	// (highly dependent on number of operators per quorum)
	UpdateStakesOfEntireOperatorSetForQuorums(
		ctx context.Context,
		operatorsPerQuorum [][]gethcommon.Address,
		quorumNumbers []byte,
	) (*gethtypes.Receipt, error)

	// UpdateStakesOfOperatorSubsetForAllQuorums is meant to be used by single operators (or teams of operators,
	// possibly running https://github.com/Layr-Labs/avs-sync) to update the stake of their own operator(s). This might
	// be needed in the case that they received a lot of new stake delegations, and want this to be reflected
	// in the AVS's registry coordinator.
	UpdateStakesOfOperatorSubsetForAllQuorums(
		ctx context.Context,
		operators []gethcommon.Address,
	) (*gethtypes.Receipt, error)

	DeregisterOperator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey regcoord.BN254G1Point,
	) (*gethtypes.Receipt, error)
}

type AvsRegistryChainWriter struct {
	serviceManagerAddr     gethcommon.Address
	registryCoordinator    *regcoord.ContractRegistryCoordinator
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry          *stakeregistry.ContractStakeRegistry
	blsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	elReader               exocontracts.EXOReader
	logger                 logging.Logger
	ethClient              eth.EthClient
	txMgr                  txmgr.TxManager
}

var _ AvsRegistryWriter = (*AvsRegistryChainWriter)(nil)

func NewAvsRegistryChainWriter(
	serviceManagerAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsApkRegistry *blsapkregistry.ContractBLSApkRegistry,
	elReader exocontracts.EXOReader,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsRegistryChainWriter, error) {
	return &AvsRegistryChainWriter{
		serviceManagerAddr:     serviceManagerAddr,
		registryCoordinator:    registryCoordinator,
		operatorStateRetriever: operatorStateRetriever,
		stakeRegistry:          stakeRegistry,
		blsApkRegistry:         blsApkRegistry,
		elReader:               elReader,
		logger:                 logger,
		ethClient:              ethClient,
		txMgr:                  txMgr,
	}, nil
}

func BuildAvsRegistryChainWriter(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	logger logging.Logger,
	ethClient eth.EthClient,
	txMgr txmgr.TxManager,
) (*AvsRegistryChainWriter, error) {
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, err
	}
	operatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
	}
	serviceManagerAddr, err := registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsApkRegistryAddr, err := registryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(blsApkRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistry, err := stakeregistry.NewContractStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, err
	}
	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	elReader, err := exocontracts.BuildExoChainReader(delegationManagerAddr, ethClient, logger)
	if err != nil {
		return nil, err
	}
	return NewAvsRegistryChainWriter(
		serviceManagerAddr,
		registryCoordinator,
		operatorStateRetriever,
		stakeRegistry,
		blsApkRegistry,
		elReader,
		logger,
		ethClient,
		txMgr,
	)
}

// TODO(samlaf): clean up this function
func (w *AvsRegistryChainWriter) RegisterOperatorInQuorumWithAVSRegistryCoordinator(
	ctx context.Context,
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	operatorToAvsRegistrationSigSalt [32]byte,
	operatorToAvsRegistrationSigExpiry *big.Int,
	blsKeyPair *bls.KeyPair,
	quorumNumbers []byte,
	socket string,
) (*gethtypes.Receipt, error) {
	w.logger.Info("registering operator with the AVS's registry coordinator")
	// params to register bls pubkey with bls apk registry
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return nil, err
	}
	signedMsg := utils.ConvertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(utils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{},
		operatorAddr,
		w.serviceManagerAddr,
		operatorToAvsRegistrationSigSalt,
		operatorToAvsRegistrationSigExpiry,
	)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	// this is annoying, and not sure why its needed, but seems like some historical baggage
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      operatorToAvsRegistrationSigSalt,
		Expiry:    operatorToAvsRegistrationSigExpiry,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		quorumNumbers,
		socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("registered operator with the AVS's registry coordinator")
	return receipt, nil
}

func (w *AvsRegistryChainWriter) UpdateStakesOfEntireOperatorSetForQuorums(
	ctx context.Context,
	operatorsPerQuorum [][]gethcommon.Address,
	quorumNumbers []byte,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes for entire operator set", "quorumNumbers", quorumNumbers)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.UpdateOperatorsForQuorum(noSendTxOpts, operatorsPerQuorum, quorumNumbers)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("updated stakes for entire operator set", "quorumNumbers", quorumNumbers)
	return receipt, nil

}

func (w *AvsRegistryChainWriter) UpdateStakesOfOperatorSubsetForAllQuorums(
	ctx context.Context,
	operators []gethcommon.Address,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes of operator subset for all quorums", "operators", operators)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.UpdateOperators(noSendTxOpts, operators)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("updated stakes of operator subset for all quorums", "operators", operators)
	return receipt, nil

}

func (w *AvsRegistryChainWriter) DeregisterOperator(
	ctx context.Context,
	quorumNumbers []byte,
	pubkey regcoord.BN254G1Point,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.registryCoordinator.DeregisterOperator(noSendTxOpts, quorumNumbers)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Info("deregistered operator with the AVS's registry coordinator")
	return receipt, nil
}
