package chainio

//go:generate mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/ExocoreNetwork/exocore-sdk/chainio/clients/avsregistry AvsRegistryReader
//go:generate mockgen -destination=./mocks/avsRegistryContractsSubscriber.go -package=mocks github.com/ExocoreNetwork/exocore-sdk/chainio/clients/avsregistry AvsRegistrySubscriber
//go:generate mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks github.com/ExocoreNetwork/exocore-sdk/chainio/clients/avsregistry AvsRegistryWriter
//go:generate mockgen -destination=./mocks/ethclient.go -package=mocks github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth EthClient
//go:generate mockgen -destination=./mocks/eventSubscription.go -package=mocks github.com/ethereum/go-ethereum/event Subscription
