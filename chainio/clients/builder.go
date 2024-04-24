package main

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/avsregistry"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exoprecompile"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type BuildAllConfig struct {
	EthHttpUrl                 string
	EthWsUrl                   string
	depositAddr                string
	avsAddr                    string
	avsTaskAddr                string
	blsAddr                    string
	RegistryCoordinatorAddr    string
	OperatorStateRetrieverAddr string
	AvsName                    string
}

type Clients struct {
	AvsRegistryChainReader     *avsregistry.AvsRegistryChainReader
	AvsRegistryChainSubscriber *avsregistry.AvsRegistryChainSubscriber
	AvsRegistryChainWriter     *avsregistry.AvsRegistryChainWriter
	EXOChainReader             *exocontracts.EXOChainReader
	EXOChainWriter             *exocontracts.EXOChainWriter
	EthHttpClient              *eth.Client
	EthWsClient                *eth.Client
}

func BuildAll(
	config BuildAllConfig,
	signerAddr gethcommon.Address,
	signerFn signerv2.SignerFn,
	logger logging.Logger,
) (*Clients, error) {
	config.validate(logger)

	// creating two types of Eth clients: HTTP and WS
	ethHttpClient, err := eth.NewClient(config.EthHttpUrl)
	if err != nil {
		logger.Error("Failed to create Eth Http client", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(config.EthWsUrl)
	if err != nil {
		logger.Error("Failed to create Eth WS client", "err", err)
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(ethHttpClient, logger, signerFn, signerAddr)
	// creating Exo clients: Reader, Writer and Subscriber
	exoChainReader, exoChainWriter, err := config.buildExoClients(
		ethHttpClient,
		txMgr,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create Exo Reader, Writer and Subscriber", "err", err)
		return nil, err
	}

	// creating AVS clients: Reader and Writer
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryChainWriter, err := config.buildAvsClients(
		exoChainReader,
		ethHttpClient,
		ethWsClient,
		txMgr,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVS Registry Reader and Writer", "err", err)
		return nil, err
	}

	return &Clients{
		EXOChainReader:             exoChainReader,
		EXOChainWriter:             exoChainWriter,
		AvsRegistryChainReader:     avsRegistryChainReader,
		AvsRegistryChainSubscriber: avsRegistryChainSubscriber,
		AvsRegistryChainWriter:     avsRegistryChainWriter,
		EthHttpClient:              ethHttpClient,
		EthWsClient:                ethWsClient,
	}, nil

}

func (config *BuildAllConfig) buildExoClients(
	ethHttpClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*exocontracts.EXOChainReader, *exocontracts.EXOChainWriter, error) {
	exoContractBindings, err := chainioutils.NewExocoreContractBindings(

		gethcommon.HexToAddress(config.depositAddr),
		gethcommon.HexToAddress(config.avsAddr),
		gethcommon.HexToAddress(config.avsTaskAddr),
		gethcommon.HexToAddress(config.blsAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ExocoreContractBindings", "err", err)
		return nil, nil, err
	}

	// get the Reader for the Exo contracts
	exoChainReader := exocontracts.NewExoChainReader(
		*exoContractBindings.DepositManager,
		*exoContractBindings.AVSManager,
		*exoContractBindings.AvsTaskManager,
		*exoContractBindings.BlsManager,
		logger,
		ethHttpClient,
	)

	elChainWriter := exocontracts.NewELChainWriter(
		*exoContractBindings.DepositManager,
		*exoContractBindings.AVSManager,
		*exoContractBindings.AvsTaskManager,
		*exoContractBindings.BlsManager,
		exoChainReader,
		ethHttpClient,
		logger,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create EXOChainWriter", "err", err)
		return nil, nil, err
	}

	return exoChainReader, elChainWriter, nil
}

func (config *BuildAllConfig) buildAvsClients(
	elReader exocontracts.EXOReader,
	ethHttpClient eth.EthClient,
	ethWsClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*avsregistry.AvsRegistryChainReader, *avsregistry.AvsRegistryChainSubscriber, *avsregistry.AvsRegistryChainWriter, error) {

	avsRegistryContractBindings, err := chainioutils.NewAVSRegistryContractBindings(
		gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
		gethcommon.HexToAddress(config.OperatorStateRetrieverAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryContractBindings", "err", err)
		return nil, nil, nil, err
	}

	avsRegistryChainReader := avsregistry.NewAvsRegistryChainReader(
		avsRegistryContractBindings.RegistryCoordinatorAddr,
		avsRegistryContractBindings.BlsApkRegistryAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		logger,
		ethHttpClient,
	)

	avsRegistryChainWriter, err := avsregistry.NewAvsRegistryChainWriter(
		avsRegistryContractBindings.ServiceManagerAddr,
		avsRegistryContractBindings.RegistryCoordinator,
		avsRegistryContractBindings.OperatorStateRetriever,
		avsRegistryContractBindings.StakeRegistry,
		avsRegistryContractBindings.BlsApkRegistry,
		elReader,
		logger,
		ethHttpClient,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create AVSRegistryChainWriter", "err", err)
		return nil, nil, nil, err
	}

	// get the Subscriber for Avs Registry contracts
	// note that the subscriber needs a ws connection instead of http
	avsRegistrySubscriber, err := avsregistry.BuildAvsRegistryChainSubscriber(
		avsRegistryContractBindings.BlsApkRegistryAddr,
		ethWsClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ELChainSubscriber", "err", err)
		return nil, nil, nil, err
	}

	return avsRegistryChainReader, avsRegistrySubscriber, avsRegistryChainWriter, nil
}

// Very basic validation that makes sure all fields are nonempty
// we might eventually want more sophisticated validation, based on regexp,
// or use something like https://json-schema.org/ (?)
func (config *BuildAllConfig) validate(logger logging.Logger) {
	if config.EthHttpUrl == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing eth http url")
	}
	if config.EthWsUrl == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing eth ws url")
	}
	if config.RegistryCoordinatorAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls registry coordinator address")
	}
	if config.OperatorStateRetrieverAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls operator state retriever address")
	}
	if config.AvsName == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing avs name")
	}
}
