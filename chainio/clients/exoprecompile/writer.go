package exocontracts

import (
	"context"
	"errors"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	avs "github.com/ExocoreNetwork/exocore-sdk/contracts/out/avs"
	avsTask "github.com/ExocoreNetwork/exocore-sdk/contracts/out/avsTask"
	bls "github.com/ExocoreNetwork/exocore-sdk/contracts/out/bls"
	deposit "github.com/ExocoreNetwork/exocore-sdk/contracts/out/deposit"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EXOWriter interface {
	DepositTo(ctx context.Context, clientChainLzID uint16, usdtAddress []byte, stakerAddr []byte, opAmount *big.Int) (*gethtypes.Receipt, error)
	AVSAction(ctx context.Context, avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*gethtypes.Receipt, error)
	RegisterAVSTask(ctx context.Context, TaskContractAddress string, Name string, MetaInfo string) (*gethtypes.Receipt, error)
}

type EXOChainWriter struct {
	depositManager deposit.Contractdeposit
	avsManager     avs.Contractavs
	avsTaskManager avsTask.ContractavsTask
	blsManager     bls.Contractbls
	exoChainReader EXOReader
	ethClient      eth.EthClient
	logger         logging.Logger
	txMgr          txmgr.TxManager
}

var _ EXOWriter = (*EXOChainWriter)(nil)

func NewELChainWriter(
	depositManager deposit.Contractdeposit,
	avsManager avs.Contractavs,
	avsTaskManager avsTask.ContractavsTask,
	blsManager bls.Contractbls,
	exoChainReader EXOReader,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *EXOChainWriter {
	return &EXOChainWriter{
		depositManager: depositManager,
		avsManager:     avsManager,
		avsTaskManager: avsTaskManager,
		blsManager:     blsManager,
		exoChainReader: exoChainReader,
		logger:         logger,
		ethClient:      ethClient,
		txMgr:          txMgr,
	}
}

func BuildELChainWriter(
	depositAddr gethcommon.Address,
	avsAddr gethcommon.Address,
	avsTaskAddr gethcommon.Address,
	blsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*EXOChainWriter, error) {
	exoContractBindings, err := chainioutils.NewExocoreContractBindings(
		depositAddr,
		avsAddr,
		avsTaskAddr,
		blsAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewExoChainReader(
		*exoContractBindings.DepositManager,
		*exoContractBindings.AVSManager,
		*exoContractBindings.AvsTaskManager,
		*exoContractBindings.BlsManager,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		*exoContractBindings.DepositManager,
		*exoContractBindings.AVSManager,
		*exoContractBindings.AvsTaskManager,
		*exoContractBindings.BlsManager,
		elChainReader,
		ethClient,
		logger,
		txMgr,
	), nil
}

func (w *EXOChainWriter) DepositTo(ctx context.Context, clientChainLzID uint16, usdtAddress []byte, stakerAddr []byte, opAmount *big.Int) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.depositManager.DepositTo(noSendTxOpts, uint16(clientChainLzID), usdtAddress, stakerAddr, opAmount)
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
func (w *EXOChainWriter) AVSAction(ctx context.Context, avsName string, avsAddress string, operatorAddress string, action uint64, avsOwnerAddress string, assetID string) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.AVSAction(noSendTxOpts, avsName, avsAddress, operatorAddress, action, avsOwnerAddress, assetID)
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

func (w *EXOChainWriter) RegisterAVSTask(ctx context.Context, TaskContractAddress string, Name string, MetaInfo string) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsTaskManager.RegisterAVSTask(noSendTxOpts, TaskContractAddress, Name, MetaInfo)
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
