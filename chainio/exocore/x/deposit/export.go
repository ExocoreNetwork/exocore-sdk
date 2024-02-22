package deposit

import sdk "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"

type Client interface {
	sdk.Module
	QueryParams(req QueryParamsRequest) (QueryParamsResponse, error)
	UpdateParams(req MsgUpdateParams) (string, error)
}
