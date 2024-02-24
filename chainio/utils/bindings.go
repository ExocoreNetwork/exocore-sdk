package utils

import (
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	blsapkregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSApkRegistry"
	delegationmanager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/DelegationManager"
	opstateretriever "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/OperatorStateRetriever"
	regcoordinator "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/StakeRegistry"
)

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type ExocoreContractBindings struct {
	SlasherAddr           gethcommon.Address
	StrategyManagerAddr   gethcommon.Address
	DelegationManagerAddr gethcommon.Address
	DelegationManager     *delegationmanager.ContractDelegationManager
}

func NewExocoreContractBindings(
	delegationManagerAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*ExocoreContractBindings, error) {
	contractDelegationManager, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch DelegationManager contract", "err", err)
		return nil, err
	}

	slasherAddr, err := contractDelegationManager.Slasher(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch Slasher address", "err", err)
		return nil, err
	}

	strategyManagerAddr, err := contractDelegationManager.StrategyManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch StrategyManager address", "err", err)
		return nil, err
	}

	return &ExocoreContractBindings{
		SlasherAddr:           slasherAddr,
		StrategyManagerAddr:   strategyManagerAddr,
		DelegationManagerAddr: delegationManagerAddr,
		DelegationManager:     contractDelegationManager,
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
