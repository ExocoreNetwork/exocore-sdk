package restaking_assets_manage

import (
	"github.com/ExocoreNetwork/exocore-sdk/common/codec"
	"github.com/ExocoreNetwork/exocore-sdk/common/codec/types"
	sdk "github.com/ExocoreNetwork/exocore-sdk/types"
)

type assetClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func (a assetClient) Name() string {
	return ModuleName
}

func (a assetClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}

func (a assetClient) SubmitProposal(request SubmitProposalRequest, baseTx sdk.BaseTx) (uint64, sdk.ResultTx, sdk.Error) {
	//TODO implement me
	panic("implement me")
}

func (a assetClient) QueryProposal(proposalId uint64) (QueryProposalResp, sdk.Error) {
	//TODO implement me
	panic("implement me")
}

func (a assetClient) QueryParams(paramsType string) (QueryParamsResp, sdk.Error) {
	//TODO implement me
	panic("implement me")
}

func NewClient(baseClient sdk.BaseClient, marshaler codec.Marshaler) Client {
	return assetClient{
		BaseClient: baseClient,
		Marshaler:  marshaler,
	}
}
