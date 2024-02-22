package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type Response interface {
	Convert() interface{}
}

type SplitAble interface {
	Len() int
	Sub(begin, end int) SplitAble
}

type Module interface {
	Name() string
	RegisterInterfaceTypes(registry codectypes.InterfaceRegistry)
}
