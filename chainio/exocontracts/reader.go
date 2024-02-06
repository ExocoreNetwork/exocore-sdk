package exocontracts

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/types"
	exoabi "github.com/ExocoreNetwork/exocore-sdk/types/abi"
)

type EXOReader interface {
	IsOperatorRegistered(ctx context.Context, operator types.Operator) (bool, error)

	GetOperatorDetails(ctx context.Context, operator types.Operator) (types.Operator, error)

	QueryExistingRegisteredOperatorPubKeys(
		startBlock *big.Int,
		stopBlock *big.Int,
	) ([]types.OperatorAddr, []types.OperatorPubkeys, error)

	GetOperatorPubkeyHash(ctx context.Context, operator types.Operator) ([32]byte, error)

	GetOperatorAddressFromPubkeyHash(ctx context.Context, pubkeyHash [32]byte) (gethcommon.Address, error)

	ServiceManagerCanSlashOperatorUntilBlock(
		ctx context.Context,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
	) (uint32, error)

	OperatorIsFrozen(ctx context.Context, operatorAddr gethcommon.Address) (bool, error)

	GetOperatorSharesInStrategy(
		ctx context.Context,
		operatorAddr gethcommon.Address,
		strategyAddr gethcommon.Address,
	) (*big.Int, error)
}

type ELChainReader struct {
	logger            logging.Logger
	elContractsClient clients.EXOContractsClient
	ethClient         eth.EthClient
}

// forces EthReader to implement the chainio.Reader interface
var _ EXOReader = (*ELChainReader)(nil)

func NewEXOChainReader(
	elContractsClient clients.EXOContractsClient,
	logger logging.Logger,
	ethClient eth.EthClient,
) (*ELChainReader, error) {
	return &ELChainReader{
		elContractsClient: elContractsClient,
		logger:            logger,
		ethClient:         ethClient,
	}, nil
}

func (r *ELChainReader) IsOperatorRegistered(ctx context.Context, operator types.Operator) (bool, error) {
	isOperator, err := r.elContractsClient.IsOperator(
		&bind.CallOpts{},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

func (r *ELChainReader) GetOperatorPubkeyHash(ctx context.Context, operator types.Operator) ([32]byte, error) {
	operatorPubkeyHash, err := r.elContractsClient.GetOperatorPubkeyHash(
		&bind.CallOpts{},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return [32]byte{}, err
	}

	return operatorPubkeyHash, nil
}

func (r *ELChainReader) GetOperatorAddressFromPubkeyHash(
	ctx context.Context,
	pubkeyHash [32]byte,
) (gethcommon.Address, error) {
	return r.elContractsClient.GetOperatorAddressFromPubkeyHash(
		&bind.CallOpts{},
		pubkeyHash,
	)
}

func (r *ELChainReader) GetOperatorDetails(ctx context.Context, operator types.Operator) (types.Operator, error) {
	operatorDetails, err := r.elContractsClient.OperatorDetails(
		&bind.CallOpts{},
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

func (r *ELChainReader) ServiceManagerCanSlashOperatorUntilBlock(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	serviceManagerAddr gethcommon.Address,
) (uint32, error) {
	serviceManagerCanSlashOperatorUntilBlock, err := r.elContractsClient.ContractCanSlashOperatorUntilBlock(
		&bind.CallOpts{Context: ctx}, operatorAddr, serviceManagerAddr,
	)
	if err != nil {
		return 0, err
	}
	return serviceManagerCanSlashOperatorUntilBlock, nil
}

func (r *ELChainReader) OperatorIsFrozen(ctx context.Context, operatorAddr gethcommon.Address) (bool, error) {
	operatorIsFrozen, err := r.elContractsClient.IsFrozen(&bind.CallOpts{Context: ctx}, operatorAddr)
	if err != nil {
		return false, err
	}
	return operatorIsFrozen, nil
}

func (r *ELChainReader) QueryExistingRegisteredOperatorPubKeys(
	startBlock *big.Int,
	stopBlock *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   stopBlock,
		Addresses: []gethcommon.Address{
			r.elContractsClient.GetBLSPublicKeyCompendiumContractAddress(),
		},
	}

	logs, err := r.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		r.logger.Error("Error filtering logs", "err", err)
		return nil, nil, err
	}
	r.logger.Info("logs:", "logs", logs)

	pubkeyCompendiumAbi, err := abi.JSON(bytes.NewReader(exoabi.BLSPublicKeyCompendiumAbi))
	if err != nil {
		r.logger.Error("Error getting Abi", "err", err)
		return nil, nil, err
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)

	for _, vLog := range logs {

		// get the operator address
		operatorAddr := gethcommon.HexToAddress(vLog.Topics[1].Hex())
		operatorAddresses = append(operatorAddresses, operatorAddr)

		event, err := pubkeyCompendiumAbi.Unpack("NewPubkeyRegistration", vLog.Data)
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

func (r *ELChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := r.elContractsClient.OperatorShares(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}
