package bank

import (
	commoncodec "github.com/ExocoreNetwork/exocore-sdk/common/codec"
	"github.com/ExocoreNetwork/exocore-sdk/common/codec/types"
	commoncryptocodec "github.com/ExocoreNetwork/exocore-sdk/common/crypto/codec"
	sdk "github.com/ExocoreNetwork/exocore-sdk/types"
	"github.com/ExocoreNetwork/exocore-sdk/types/auth"
)

var (
	amino     = commoncodec.NewLegacyAmino()
	ModuleCdc = commoncodec.NewAminoCodec(amino)
)

func init() {
	commoncryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// No duplicate registration
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)

	registry.RegisterImplementations(
		(*auth.Account)(nil),
		&auth.BaseAccount{},
	)

	registry.RegisterImplementations(
		(*auth.Account)(nil),
		&auth.EthAccount{},
	)
}
