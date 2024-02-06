package clients

import (
	"math/big"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	blspubkeycompendium "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSPublicKeyCompendium"
	contractDelegationManager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/DelegationManager"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type EXOContractsClient interface {

	// Write methods

	RegisterAsOperator(
		opts *bind.TransactOpts,
		registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
		metadataURI string,
	) (*gethTypes.Transaction, error)

	ModifyOperatorDetails(
		opts *bind.TransactOpts,
		registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
	) (*gethTypes.Transaction, error)

	UpdateOperatorMetadataURI(
		opts *bind.TransactOpts,
		metadataURI string,
	) (*gethTypes.Transaction, error)

	OptIntoSlashing(
		opts *bind.TransactOpts,
		avsServiceManagerAddr common.Address,
	) (*gethTypes.Transaction, error)

	RegisterBLSPublicKey(
		opts *bind.TransactOpts,
		signedMessageHash blspubkeycompendium.BN254G1Point,
		pubkeyG1 blspubkeycompendium.BN254G1Point,
		pubkeyG2 blspubkeycompendium.BN254G2Point,
	) (*gethTypes.Transaction, error)

	WatchNewPubkeyRegistration(
		opts *bind.WatchOpts,
		sink chan<- *blspubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration,
		operator []common.Address,
	) (event.Subscription, error)

	// Read methods

	GetStrategyManagerContractAddress() (common.Address, error)

	GetBLSPublicKeyCompendiumContractAddress() common.Address

	IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error)

	OperatorDetails(
		opts *bind.CallOpts,
		operator common.Address,
	) (contractDelegationManager.IDelegationManagerOperatorDetails, error)

	ContractCanSlashOperatorUntilBlock(
		opts *bind.CallOpts,
		operator common.Address,
		serviceContract common.Address,
	) (uint32, error)

	IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error)

	GetOperatorPubkeyHash(
		opts *bind.CallOpts,
		operatorAddr common.Address,
	) ([32]byte, error)

	GetOperatorAddressFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	OperatorShares(opts *bind.CallOpts, operatorAddr common.Address, strategyAddr common.Address) (*big.Int, error)
}

// EXOContractsChainClient is really just a wrapper around the go bindings that has a proper interface, so that we can
// mock it in tests.
// TODO(samlaf): should we instead just make the
type EXOContractsChainClient struct {
	exoHttpBindings *exocoreContractBindings
	// TODO(samlaf): currently we're creating a second set of bindings backed by a websocket client, so that we can
	// subscribe to events. perhaps a better way would be to make our EthClient implementation have both an http and
	// websocket client, and use the websocket client for
	// the SubscribeFilterLogs() method, and the http client for everything else...
	// Yet another option is to start using an indexer (maybe
	// https://github.com/ethereum-optimism/optimism/tree/develop/indexer) and drop websocket subscriptions altogether.
	exoWsBindings           *exocoreContractBindings
	ethHttpClient           eth.EthClient
	logger                  logging.Logger
	blsPubKeyCompendiumAddr common.Address
}

var _ EXOContractsClient = (*EXOContractsChainClient)(nil)

func NewEXOContractsChainClient(
	slasherAddr common.Address,
	blsPubKeyCompendiumAddr common.Address,
	ethHttpClient eth.EthClient,
	ethWsClient eth.EthClient,
	logger logging.Logger,
) (*EXOContractsChainClient, error) {
	exoHttpBindings, err := newExocoreContractBindings(slasherAddr, blsPubKeyCompendiumAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}
	exoWsBindings, err := newExocoreContractBindings(slasherAddr, blsPubKeyCompendiumAddr, ethWsClient, logger)
	if err != nil {
		return nil, err
	}

	return &EXOContractsChainClient{
		exoHttpBindings:         exoHttpBindings,
		exoWsBindings:           exoWsBindings,
		ethHttpClient:           ethHttpClient,
		logger:                  logger,
		blsPubKeyCompendiumAddr: blsPubKeyCompendiumAddr,
	}, nil
}

func (e *EXOContractsChainClient) RegisterAsOperator(
	opts *bind.TransactOpts,
	registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
	metadataURI string,
) (*gethTypes.Transaction, error) {
	return e.exoHttpBindings.DelegationManager.RegisterAsOperator(opts, registeringOperatorDetails, metadataURI)
}

func (e *EXOContractsChainClient) ModifyOperatorDetails(
	opts *bind.TransactOpts,
	registeringOperatorDetails contractDelegationManager.IDelegationManagerOperatorDetails,
) (*gethTypes.Transaction, error) {
	return e.exoHttpBindings.DelegationManager.ModifyOperatorDetails(opts, registeringOperatorDetails)
}

func (e *EXOContractsChainClient) UpdateOperatorMetadataURI(
	opts *bind.TransactOpts,
	metadataURI string,
) (*gethTypes.Transaction, error) {
	return e.exoHttpBindings.DelegationManager.UpdateOperatorMetadataURI(opts, metadataURI)
}

func (e *EXOContractsChainClient) OptIntoSlashing(
	opts *bind.TransactOpts,
	avsServiceManagerAddr common.Address,
) (*gethTypes.Transaction, error) {
	return e.exoHttpBindings.Slasher.OptIntoSlashing(opts, avsServiceManagerAddr)
}

func (e *EXOContractsChainClient) RegisterBLSPublicKey(
	opts *bind.TransactOpts,
	signedMessageHash blspubkeycompendium.BN254G1Point,
	pubkeyG1 blspubkeycompendium.BN254G1Point,
	pubkeyG2 blspubkeycompendium.BN254G2Point,
) (*gethTypes.Transaction, error) {
	return e.exoHttpBindings.BlsPubKeyCompendium.RegisterBLSPublicKey(opts, signedMessageHash, pubkeyG1, pubkeyG2)
}

func (e *EXOContractsChainClient) GetStrategyManagerContractAddress() (common.Address, error) {
	return e.exoHttpBindings.Slasher.StrategyManager(&bind.CallOpts{})
}

func (e *EXOContractsChainClient) WatchNewPubkeyRegistration(
	opts *bind.WatchOpts,
	sink chan<- *blspubkeycompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration,
	operator []common.Address,
) (event.Subscription, error) {
	return e.exoWsBindings.BlsPubKeyCompendium.WatchNewPubkeyRegistration(opts, sink, operator)
}

func (e *EXOContractsChainClient) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	return e.exoHttpBindings.DelegationManager.IsOperator(opts, operator)
}

func (e *EXOContractsChainClient) OperatorDetails(
	opts *bind.CallOpts,
	operator common.Address,
) (contractDelegationManager.IDelegationManagerOperatorDetails, error) {
	return e.exoHttpBindings.DelegationManager.OperatorDetails(opts, operator)
}

func (e *EXOContractsChainClient) ContractCanSlashOperatorUntilBlock(
	opts *bind.CallOpts,
	operator common.Address,
	serviceContract common.Address,
) (uint32, error) {
	return e.exoHttpBindings.Slasher.ContractCanSlashOperatorUntilBlock(opts, operator, serviceContract)
}

func (e *EXOContractsChainClient) IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error) {
	return e.exoHttpBindings.Slasher.IsFrozen(opts, staker)
}

func (e *EXOContractsChainClient) GetBLSPublicKeyCompendiumContractAddress() common.Address {
	return e.blsPubKeyCompendiumAddr
}

func (e *EXOContractsChainClient) GetOperatorPubkeyHash(
	opts *bind.CallOpts,
	operatorAddr common.Address,
) ([32]byte, error) {
	return e.exoHttpBindings.BlsPubKeyCompendium.OperatorToPubkeyHash(opts, operatorAddr)
}

func (e *EXOContractsChainClient) GetOperatorAddressFromPubkeyHash(
	opts *bind.CallOpts,
	pubkeyHash [32]byte,
) (common.Address, error) {
	return e.exoHttpBindings.BlsPubKeyCompendium.PubkeyHashToOperator(opts, pubkeyHash)
}

func (e *EXOContractsChainClient) OperatorShares(
	opts *bind.CallOpts,
	operatorAddr common.Address,
	strategyAddr common.Address,
) (*big.Int, error) {
	operatorSharesInStrategy, err := e.exoHttpBindings.DelegationManager.OperatorShares(opts, operatorAddr, strategyAddr)
	if err != nil {
		return nil, err
	}
	return operatorSharesInStrategy, nil
}
