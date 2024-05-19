package main

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exocontracts"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	chainioutils "github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type BuildAllConfig struct {
	EthHttpUrl                 string
	EthWsUrl                   string
	avsAddr                    string
	RegistryCoordinatorAddr    string
	OperatorStateRetrieverAddr string
	AvsName                    string
}

type Clients struct {
	AvsRegistryChainSubscriber *exocontracts.AvsRegistryChainSubscriber
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
	exoChainReader, exoChainWriter, avsRegistrySubscriber, err := config.buildExoClients(
		ethHttpClient,
		txMgr,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create Exo Reader, Writer and Subscriber", "err", err)
		return nil, err
	}

	return &Clients{
		EXOChainReader:             exoChainReader,
		EXOChainWriter:             exoChainWriter,
		AvsRegistryChainSubscriber: avsRegistrySubscriber,
		EthHttpClient:              ethHttpClient,
		EthWsClient:                ethWsClient,
	}, nil

}

func (config *BuildAllConfig) buildExoClients(
	ethHttpClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*exocontracts.EXOChainReader, *exocontracts.EXOChainWriter, *exocontracts.AvsRegistryChainSubscriber, error) {
	exoContractBindings, err := chainioutils.NewExocoreContractBindings(
		gethcommon.HexToAddress(config.avsAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ExocoreContractBindings", "err", err)
		return nil, nil, nil, err
	}

	// get the Reader for the Exo contracts
	exoChainReader := exocontracts.NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethHttpClient,
	)

	elChainWriter := exocontracts.NewELChainWriter(
		*exoContractBindings.AVSManager,
		exoChainReader,
		ethHttpClient,
		logger,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create EXOChainWriter", "err", err)
		return nil, nil, nil, err
	}

	avsRegistrySubscriber, err := exocontracts.BuildAvsRegistryChainSubscriber(
		exoContractBindings.AvsAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ELChainSubscriber", "err", err)
		return nil, nil, nil, err
	}
	return exoChainReader, elChainWriter, avsRegistrySubscriber, err
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
