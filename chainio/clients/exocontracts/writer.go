package exocontracts

import (
	"context"
	"errors"
	"math/big"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/types"

	delegationmanager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/DelegationManager"
)

type EXOWriter interface {
	RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	UpdateOperatorDetails(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	// DepositERC20IntoStrategy deposits ERC20 tokens into a strategy contract.
	DepositERC20IntoStrategy(
		ctx context.Context,
		strategyAddr gethcommon.Address,
		amount *big.Int,
	) (*gethtypes.Receipt, error)
}

type EXOChainWriter struct {
	delegationManager   delegationmanager.ContractDelegationManagerTransacts
	strategyManagerAddr gethcommon.Address
	exoChainReader      EXOReader
	ethClient           eth.EthClient
	logger              logging.Logger
	txMgr               txmgr.TxManager
}

func (w *EXOChainWriter) DepositERC20IntoStrategy(ctx context.Context, strategyAddr gethcommon.Address, amount *big.Int) (*gethtypes.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

var _ EXOWriter = (*EXOChainWriter)(nil)

func NewELChainWriter(
	delegationManager delegationmanager.ContractDelegationManagerTransacts,
	strategyManagerAddr gethcommon.Address,
	elChainReader EXOReader,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *EXOChainWriter {
	return &EXOChainWriter{
		delegationManager:   delegationManager,
		strategyManagerAddr: strategyManagerAddr,
		exoChainReader:      elChainReader,
		logger:              logger,
		ethClient:           ethClient,
		txMgr:               txMgr,
	}
}

func BuildELChainWriter(
	delegationManagerAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*EXOChainWriter, error) {
	elContractBindings, err := chainioutils.NewExocoreContractBindings(
		delegationManagerAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewExoChainReader(
		elContractBindings.DelegationManager,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		txMgr,
	), nil
}

// TODO(madhur): we wait for txreceipts in these functions right now, but
// this will be changed once we have a better tx manager design implemented
// see https://github.com/ExocoreNetwork/exocore-sdk/pull/75
func (w *EXOChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
	w.logger.Infof("registering operator %s to exocore", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.delegationManager.RegisterAsOperator(noSendTxOpts, opDetails, operator.MetadataUrl)
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

func (w *EXOChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (*gethtypes.Receipt, error) {

	w.logger.Infof("updating operator details of operator %s to exocore", operator.Address)
	opDetails := delegationmanager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		DelegationApprover:       gethcommon.HexToAddress(operator.DelegationApproverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(noSendTxOpts, opDetails)
	if err != nil {
		return nil, err
	}
	_, err = w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.logger.Infof("updated operator metadata URI for operator %s to exocore", operator.Address)

	tx, err = w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	w.logger.Infof("updated operator details of operator %s to exocore", operator.Address)
	return receipt, nil
}
