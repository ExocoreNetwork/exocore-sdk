package exocore

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/client"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/x/deposit"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/x/operator"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/x/restaking_assets_manage"
	"github.com/cosmos/cosmos-sdk/codec"
	atypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/log"
)

type Client struct {
	logger            log.Logger
	encodingConfig    types.EncodingConfig
	Deposit           deposit.Client
	baseclient        client.BaseClient
	AssetManageClient restaking_assets_manage.Client
	OperatorClient    operator.Client
}

func NewClient(cfg types.ClientConfig) Client {
	encodingConfig := makeEncodingConfig()

	baseClient := client.NewBaseClient(cfg, encodingConfig, log.New())

	depositClient := deposit.NewClient(baseClient)
	assetManageClient := restaking_assets_manage.NewClient(baseClient)
	operatorClient := operator.NewClient(baseClient)
	client := Client{
		logger:            nil,
		Deposit:           depositClient,
		AssetManageClient: assetManageClient,
		encodingConfig:    encodingConfig,
		baseclient:        baseClient,
		OperatorClient:    operatorClient,
	}
	client.RegisterModule(
		depositClient,
	)
	return client
}
func makeEncodingConfig() types.EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := atypes.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	encCfg := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}

func (client *Client) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
	}
}
