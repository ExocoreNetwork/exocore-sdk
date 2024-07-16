package avs

import sdk "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"

type Client interface {
	sdk.Module
	QueryAVSInfo(req QueryAVSInfoReq) (AVSInfo, error)
	RegisterAVS(req RegisterAVSReq) (string, error)
}
