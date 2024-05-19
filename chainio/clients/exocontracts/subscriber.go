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
	SubscribeToNewPubkeyRegistrations() (chan *avssub.ContractavsserviceTaskCreated, event.Subscription, error)
}

type AvsRegistryChainSubscriber struct {
	logger logging.Logger
	avssub avssub.ContractavsserviceFilterer
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ AvsRegistrySubscriber = (*AvsRegistryChainSubscriber)(nil)

func NewAvsRegistryChainSubscriber(
	avssub avssub.ContractavsserviceFilterer,
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
	avssub, err := avssub.NewContractavsserviceFilterer(avssubAddr, ethWsClient)
	if err != nil {
		logger.Error("Failed to create BLSApkRegistry contract", "err", err)
		return nil, err
	}
	return NewAvsRegistryChainSubscriber(*avssub, logger)
}

func (s *AvsRegistryChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *avssub.ContractavsserviceTaskCreated, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *avssub.ContractavsserviceTaskCreated)
	sub, err := s.avssub.WatchTaskCreated(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to NewPubkeyRegistration events", "err", err)
		return nil, nil, err
	}
	return newPubkeyRegistrationChan, sub, nil
}
