package exocore

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/client"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"
	"github.com/cosmos/cosmos-sdk/codec"
	atypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/log"
)

type Client struct {
	logger         log.Logger
	encodingConfig types.EncodingConfig
	baseclient     client.BaseClient
}

func NewClient(cfg types.ClientConfig) Client {
	encodingConfig := makeEncodingConfig()
	baseClient := client.NewBaseClient(cfg, encodingConfig, log.New())
	client := Client{
		logger:         nil,
		encodingConfig: encodingConfig,
		baseclient:     baseClient,
	}
	client.RegisterModule()
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
