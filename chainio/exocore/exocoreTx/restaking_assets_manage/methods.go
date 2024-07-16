package restaking_assets_manage

import (
	"context"
	grpc "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type assetManageClient struct {
	baseclient grpc.BaseClient
}

func NewClient(baseClient grpc.BaseClient) Client {
	return &assetManageClient{
		baseclient: baseClient,
	}
}
func (sc assetManageClient) Name() string {
	return ModuleName
}
func (sc assetManageClient) RegisterInterfaceTypes(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}
func (sc assetManageClient) QueClientChainInfoByIndex(chainIndex uint64) (QueClientChainInfoByIndexResp, error) {
	conn, err := sc.baseclient.GenConn()

	if err != nil {
		return QueClientChainInfoByIndexResp{}, err
	}

	res, err := NewQueryClient(conn).QueClientChainInfoByIndex(
		context.Background(),
		&QueryClientChainInfo{
			ChainIndex: chainIndex,
		},
	)
	if err != nil {
		return QueClientChainInfoByIndexResp{}, err
	}
	return QueClientChainInfoByIndexResp{res.ChainName, res.ChainMetaInfo, res.OriginChainId,
		res.ExoCoreChainIndex, res.FinalityNeedBlockDelay, res.LayerZeroChainId,
		res.SignatureType, res.AddressLength}, nil
}
