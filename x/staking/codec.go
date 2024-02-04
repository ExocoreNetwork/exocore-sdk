package staking

import (
	"github.com/ExocoreNetwork/exocore-sdk/common/codec"
	"github.com/ExocoreNetwork/exocore-sdk/common/codec/types"
	cryptocodec "github.com/ExocoreNetwork/exocore-sdk/common/crypto/codec"
	sdk "github.com/ExocoreNetwork/exocore-sdk/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgBeginRedelegate{},
	)
}
