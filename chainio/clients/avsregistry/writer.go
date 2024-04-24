package avsregistry

import (
	"context"
	"errors"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exoprecompile"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	blsapkregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/StakeRegistry"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type AvsRegistryWriter interface {
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
