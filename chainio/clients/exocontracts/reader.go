package exocontracts

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	avs "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EXOReader interface {
	//	PublicKey          []byte
	//Name               string
	//TotalRewardsEarned *big.Int
	//IsRegistered       bool
	Operators(opts *bind.CallOpts, operatorAddress gethcommon.Address) (Operator, error)
}
type Operator struct {
	PublicKey          []byte
	Name               string
	TotalRewardsEarned *big.Int
	IsRegistered       bool
}
type EXOChainReader struct {
	logger     logging.Logger
	avsManager avs.Contractavsservice
	ethClient  eth.EthClient
}

// forces EthReader to implement the chainio.Reader interface
var _ EXOReader = (*EXOChainReader)(nil)

func NewExoChainReader(
	avsManager avs.Contractavsservice,
	logger logging.Logger,
	ethClient eth.EthClient,
) *EXOChainReader {
	return &EXOChainReader{
		avsManager: avsManager,
		logger:     logger,
		ethClient:  ethClient,
	}
}

func BuildExoChainReader(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*EXOChainReader, error) {
	exoContractBindings, err := chainioutils.NewExocoreContractBindings(
		avsAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethClient,
	), nil
}

func (r *EXOChainReader) Operators(opts *bind.CallOpts, operatorAddress gethcommon.Address) (Operator, error) {

	operator, err := r.avsManager.Operators(
		opts,
		operatorAddress,
	)
	if err != nil {
		r.logger.Error("Failed to get operator ", "err", err)
		return Operator{}, err
	}
	return operator, nil
}
