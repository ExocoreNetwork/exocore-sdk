package exocontracts

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	avs "github.com/ExocoreNetwork/exocore-sdk/contracts/out/avs"
	avsTask "github.com/ExocoreNetwork/exocore-sdk/contracts/out/avsTask"
	bls "github.com/ExocoreNetwork/exocore-sdk/contracts/out/bls"
	deposit "github.com/ExocoreNetwork/exocore-sdk/contracts/out/deposit"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type EXOReader interface{}
type EXOChainReader struct {
	logger         logging.Logger
	depositManager deposit.Contractdeposit
	avsManager     avs.Contractavs
	avsTaskManager avsTask.ContractavsTask
	blsManager     bls.Contractbls
	ethClient      eth.EthClient
}

// forces EthReader to implement the chainio.Reader interface
var _ EXOReader = (*EXOChainReader)(nil)

func NewExoChainReader(
	depositManager deposit.Contractdeposit,
	avsManager avs.Contractavs,
	avsTaskManager avsTask.ContractavsTask,
	blsManager bls.Contractbls,
	logger logging.Logger,
	ethClient eth.EthClient,
) *EXOChainReader {
	return &EXOChainReader{
		depositManager: depositManager,
		avsManager:     avsManager,
		avsTaskManager: avsTaskManager,
		blsManager:     blsManager,
		logger:         logger,
		ethClient:      ethClient,
	}
}

func BuildExoChainReader(
	depositAddr gethcommon.Address,
	avsAddr gethcommon.Address,
	avsTaskAddr gethcommon.Address,
	blsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*EXOChainReader, error) {
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
	return NewExoChainReader(
		*exoContractBindings.DepositManager,
		*exoContractBindings.AVSManager,
		*exoContractBindings.AvsTaskManager,
		*exoContractBindings.BlsManager,
		logger,
		ethClient,
	), nil
}
