package types

import (
	"github.com/ExocoreNetwork/exocore-sdk/common/codec"
	"github.com/ExocoreNetwork/exocore-sdk/common/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          TxConfig
	Amino             *codec.LegacyAmino
}
