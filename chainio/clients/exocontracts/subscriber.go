package exocontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	avssub "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
)

type AvsRegistrySubscriber interface {
	SubscribeToNewTasks(newTaskCreatedChan chan *avssub.ContractavsserviceTaskCreated) event.Subscription
}

type AvsRegistryChainSubscriber struct {
	logger logging.Logger
	avssub avssub.Contractavsservice
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ AvsRegistrySubscriber = (*AvsRegistryChainSubscriber)(nil)

func NewAvsRegistryChainSubscriber(
	avssub avssub.Contractavsservice,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	return &AvsRegistryChainSubscriber{
		logger: logger,
		avssub: avssub,
	}, nil
}

func BuildAvsRegistryChainSubscriber(
	avssubAddr common.Address,
	ethWsClient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	avssub, err := avssub.NewContractavsservice(avssubAddr, ethWsClient)
	if err != nil {
		logger.Error("Failed to create BLSApkRegistry contract", "err", err)
		return nil, err
	}
	return NewAvsRegistryChainSubscriber(*avssub, logger)
}

func (s *AvsRegistryChainSubscriber) SubscribeToNewTasks(newTaskCreatedChan chan *avssub.ContractavsserviceTaskCreated) event.Subscription {
	sub, err := s.avssub.WatchTaskCreated(
		&bind.WatchOpts{}, newTaskCreatedChan,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new  tasks", "err", err)
	}
	s.logger.Infof("Subscribed to new TaskManager tasks")
	return sub
}
