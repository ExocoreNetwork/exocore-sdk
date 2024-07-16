package operator

import sdk "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"

type Client interface {
	sdk.Module
	GetOperatorInfo(req GetOperatorInfoReq) (OperatorInfo, error)
	RegisterOperator(req RegisterOperatorReq) (string, error)
}
