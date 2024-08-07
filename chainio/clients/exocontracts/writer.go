package exocontracts

import (
	"context"
	"errors"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	avs "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type EXOWriter interface {
	RegisterAVSToExocore(
		ctx context.Context,
		avsName string,
		minStakeAmount uint64,
		taskAddr gethcommon.Address,
		slashAddr gethcommon.Address,
		rewardAddr gethcommon.Address,
		avsOwnerAddress []string,
		assetIds []string,
		avsUnbondingPeriod uint64,
		minSelfDelegation uint64,
		epochIdentifier string,
		params []uint64,
	) (*gethtypes.Receipt, error)
}

type EXOChainWriter struct {
	avsManager     avs.Contractavsservice
	exoChainReader EXOReader
	ethClient      eth.EthClient
	logger         logging.Logger
	txMgr          txmgr.TxManager
}

var _ EXOWriter = (*EXOChainWriter)(nil)

func NewELChainWriter(
	avsManager avs.Contractavsservice,
	exoChainReader EXOReader,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *EXOChainWriter {
	return &EXOChainWriter{
		avsManager:     avsManager,
		exoChainReader: exoChainReader,
		logger:         logger,
		ethClient:      ethClient,
		txMgr:          txMgr,
	}
}

func BuildELChainWriter(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*EXOChainWriter, error) {
	exoContractBindings, err := chainioutils.NewExocoreContractBindings(
		avsAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		*exoContractBindings.AVSManager,
		elChainReader,
		ethClient,
		logger,
		txMgr,
	), nil
}

func (w *EXOChainWriter) RegisterAVSToExocore(
	ctx context.Context,
	avsName string,
	minStakeAmount uint64,
	taskAddr gethcommon.Address,
	slashAddr gethcommon.Address,
	rewardAddr gethcommon.Address,
	avsOwnerAddress []string,
	assetIds []string,
	avsUnbondingPeriod uint64,
	minSelfDelegation uint64,
	epochIdentifier string,
	params []uint64,
) (*gethtypes.Receipt, error) {

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterAVS(noSendTxOpts,
		avsName,
		minStakeAmount,
		taskAddr,
		slashAddr,
		rewardAddr,
		avsOwnerAddress,
		assetIds,
		avsUnbondingPeriod,
		minSelfDelegation,
		epochIdentifier,
		params)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	return receipt, nil
}
