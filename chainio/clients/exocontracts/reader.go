package exocontracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/types"

	delegationmanager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/DelegationManager"
)

type EXOReader interface {
	IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error)

	GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error)

	ServiceManagerCanSlashOperatorUntilBlock(
		opts *bind.CallOpts,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
	) (uint32, error)

	OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error)

	GetOperatorSharesInStrategy(
		opts *bind.CallOpts,
		operatorAddr gethcommon.Address,
		strategyAddr gethcommon.Address,
	) (*big.Int, error)

	CalculateDelegationApprovalDigestHash(
		opts *bind.CallOpts, operator gethcommon.Address, avs gethcommon.Address, salt [32]byte, expiry *big.Int,
	) ([32]byte, error)
}

type EXOChainReader struct {
	logger            logging.Logger
	delegationManager delegationmanager.ContractDelegationManagerCalls
	ethClient         eth.EthClient
}

func (r *EXOChainReader) ServiceManagerCanSlashOperatorUntilBlock(opts *bind.CallOpts, operatorAddr gethcommon.Address, serviceManagerAddr gethcommon.Address) (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (r *EXOChainReader) OperatorIsFrozen(opts *bind.CallOpts, operatorAddr gethcommon.Address) (bool, error) {
	//TODO implement me
	panic("implement me")
}

// forces EthReader to implement the chainio.Reader interface
var _ EXOReader = (*EXOChainReader)(nil)

func NewExoChainReader(
	delegationManager delegationmanager.ContractDelegationManagerCalls,
	logger logging.Logger,
	ethClient eth.EthClient,
) *EXOChainReader {
	return &EXOChainReader{
		delegationManager: delegationManager,
		logger:            logger,
		ethClient:         ethClient,
	}
}

func BuildExoChainReader(
	delegationManagerAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*EXOChainReader, error) {
	exoContractBindings, err := chainioutils.NewExocoreContractBindings(
		delegationManagerAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewExoChainReader(
		exoContractBindings.DelegationManager,
		logger,
		ethClient,
	), nil
}

func (r *EXOChainReader) IsOperatorRegistered(opts *bind.CallOpts, operator types.Operator) (bool, error) {
	isOperator, err := r.delegationManager.IsOperator(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

func (r *EXOChainReader) GetOperatorDetails(opts *bind.CallOpts, operator types.Operator) (types.Operator, error) {
	operatorDetails, err := r.delegationManager.OperatorDetails(
		opts,
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return types.Operator{}, err
	}

	return types.Operator{
		Address:                   operator.Address,
		EarningsReceiverAddress:   operatorDetails.EarningsReceiver.Hex(),
		StakerOptOutWindowBlocks:  operatorDetails.StakerOptOutWindowBlocks,
		DelegationApproverAddress: operatorDetails.DelegationApprover.Hex(),
	}, nil
}

func (r *EXOChainReader) GetOperatorSharesInStrategy(
	opts *bind.CallOpts,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := r.delegationManager.OperatorShares(
		opts,
		operatorAddr,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}

func (r *EXOChainReader) CalculateDelegationApprovalDigestHash(
	opts *bind.CallOpts, operator gethcommon.Address, avs gethcommon.Address, salt [32]byte, expiry *big.Int,
) ([32]byte, error) {
	return r.delegationManager.CalculateOperatorAVSRegistrationDigestHash(
		opts, operator, avs, salt, expiry,
	)
}
