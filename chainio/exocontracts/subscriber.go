package exocontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients"
	pubkeycompendium "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSPublicKeyCompendium"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
)

type EXOSubscriber interface {
	SubscribeToNewPubkeyRegistrations() (chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error)
}

type EXOChainSubscriber struct {
	logger            logging.Logger
	elContractsClient clients.EXOContractsClient
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ EXOSubscriber = (*EXOChainSubscriber)(nil)

func NewEXOChainSubscriber(
	elContractsClient clients.EXOContractsClient,
	logger logging.Logger,
) (*EXOChainSubscriber, error) {

	return &EXOChainSubscriber{
		logger:            logger,
		elContractsClient: elContractsClient,
	}, nil
}

func (s *EXOChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *pubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
	sub, err := s.elContractsClient.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to NewPubkeyRegistration events", "err", err)
		return nil, nil, err
	}
	return newPubkeyRegistrationChan, sub, nil
}
