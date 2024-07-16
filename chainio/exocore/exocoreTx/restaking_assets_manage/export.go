package restaking_assets_manage

import sdk "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"

type Client interface {
	sdk.Module
	QueClientChainInfoByIndex(chainIndex uint64) (QueClientChainInfoByIndexResp, error)
}

type QueClientChainInfoByIndexResp struct {
	ChainName              string `json:"ChainName,omitempty"`
	ChainMetaInfo          string `json:"ChainMetaInfo,omitempty"`
	OriginChainId          uint64 `json:"OriginChainId,omitempty"`
	ExoCoreChainIndex      uint64 `json:"ExoCoreChainIndex,omitempty"`
	FinalityNeedBlockDelay uint64 `json:"FinalityNeedBlockDelay,omitempty"`
	LayerZeroChainId       uint64 `json:"LayerZeroChainId,omitempty"`
	SignatureType          string `json:"SignatureType,omitempty"`
	AddressLength          uint32 `json:"AddressLength,omitempty"`
}
