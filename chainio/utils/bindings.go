package utils

import (
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	blsapkregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/OperatorStateRetriever"
	regcoordinator "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/StakeRegistry"
	avs "github.com/ExocoreNetwork/exocore-sdk/contracts/out/avs"
	avsTask "github.com/ExocoreNetwork/exocore-sdk/contracts/out/avsTask"
	bls "github.com/ExocoreNetwork/exocore-sdk/contracts/out/bls"
	deposit "github.com/ExocoreNetwork/exocore-sdk/contracts/out/deposit"
)

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type ExocoreContractBindings struct {
	DepositAddr    gethcommon.Address
	DepositManager *deposit.Contractdeposit
	AvsAddr        gethcommon.Address
	AVSManager     *avs.Contractavs
	AvsTaskAddr    gethcommon.Address
	AvsTaskManager *avsTask.ContractavsTask
	BlsAddr        gethcommon.Address
	BlsManager     *bls.Contractbls
}

func NewExocoreContractBindings(
	depositAddr gethcommon.Address,
	avsAddr gethcommon.Address,
	avsTaskAddr gethcommon.Address,
	blsAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*ExocoreContractBindings, error) {
	contractDepositManager, err := deposit.NewContractdeposit(depositAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch Deposit contract", "err", err)
		return nil, err
	}

	contractAvsManager, err := avs.NewContractavs(avsAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch Avs contract", "err", err)
		return nil, err
	}

	contractAvsTaskManager, err := avsTask.NewContractavsTask(avsTaskAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch AvsTask contract", "err", err)
		return nil, err
	}

	contractBlsManager, err := bls.NewContractbls(blsAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch bls contract", "err", err)
		return nil, err
	}
	return &ExocoreContractBindings{
		DepositAddr:    depositAddr,
		DepositManager: contractDepositManager,
		AvsAddr:        avsAddr,
		AVSManager:     contractAvsManager,
		AvsTaskAddr:    avsTaskAddr,
		AvsTaskManager: contractAvsTaskManager,
		BlsAddr:        blsAddr,
		BlsManager:     contractBlsManager,
	}, nil
}

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type AvsRegistryContractBindings struct {
	// contract addresses
	ServiceManagerAddr         gethcommon.Address
	RegistryCoordinatorAddr    gethcommon.Address
	StakeRegistryAddr          gethcommon.Address
	BlsApkRegistryAddr         gethcommon.Address
	OperatorStateRetrieverAddr gethcommon.Address
	// contract bindings
	ServiceManager         *servicemanager.ContractServiceManagerBase
	RegistryCoordinator    *regcoordinator.ContractRegistryCoordinator
	StakeRegistry          *stakeregistry.ContractStakeRegistry
	BlsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	OperatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
}

func NewAVSRegistryContractBindings(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryContractBindings, error) {
	contractBlsRegistryCoordinator, err := regcoordinator.NewContractRegistryCoordinator(
		registryCoordinatorAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSRegistryCoordinator contract", "err", err)
		return nil, err
	}

	serviceManagerAddr, err := contractBlsRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch ServiceManager address", "err", err)
		return nil, err
	}
	contractServiceManager, err := servicemanager.NewContractServiceManagerBase(
		serviceManagerAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch ServiceManager contract", "err", err)
		return nil, err
	}

	stakeregistryAddr, err := contractBlsRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch StakeRegistry address", "err", err)
		return nil, err
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(
		stakeregistryAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch StakeRegistry contract", "err", err)
		return nil, err
	}

	blsApkRegistryAddr, err := contractBlsRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch BLSPubkeyRegistry address", "err", err)
		return nil, err
	}
	contractBlsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(
		blsApkRegistryAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSPubkeyRegistry contract", "err", err)
		return nil, err
	}

	contractOperatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch OperatorStateRetriever contract", "err", err)
		return nil, err
	}

	return &AvsRegistryContractBindings{
		ServiceManagerAddr:         serviceManagerAddr,
		RegistryCoordinatorAddr:    registryCoordinatorAddr,
		StakeRegistryAddr:          stakeregistryAddr,
		BlsApkRegistryAddr:         blsApkRegistryAddr,
		OperatorStateRetrieverAddr: operatorStateRetrieverAddr,
		ServiceManager:             contractServiceManager,
		RegistryCoordinator:        contractBlsRegistryCoordinator,
		StakeRegistry:              contractStakeRegistry,
		BlsApkRegistry:             contractBlsApkRegistry,
		OperatorStateRetriever:     contractOperatorStateRetriever,
	}, nil
}
