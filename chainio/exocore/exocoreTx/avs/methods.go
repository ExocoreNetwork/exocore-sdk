package avs

import (
	"context"
	grpc "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/client"
	util "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type avsClient struct {
	baseclient grpc.BaseClient
}

func (sc avsClient) RegisterInterfaceTypes(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}
func (sc avsClient) Name() string {
	return ModuleName
}

func NewClient(baseClient grpc.BaseClient) *avsClient {
	return &avsClient{
		baseclient: baseClient,
	}
}

func (sc avsClient) GetOperatorInfo(req QueryAVSInfoReq) (QueryAVSInfoResponse, error) {
	conn, err := sc.baseclient.GenConn()

	if err != nil {
		return QueryAVSInfoResponse{}, err
	}

	res, err := NewQueryClient(conn).QueryAVSInfo(
		context.Background(),
		&req,
	)

	if err != nil {
		return QueryAVSInfoResponse{}, err
	}
	return *res, nil
}

func (sc avsClient) RegisterOperator(req RegisterAVSReq) (string, error) {

	accountAddress := sdk.AccAddress(sc.baseclient.GenCfg().Priv.PubKey().Address().Bytes())
	msg := &RegisterAVSReq{
		FromAddress: util.GetBech32Addr("exo", accountAddress),

		Info: &AVSInfo{
			Name:               req.Info.Name,
			AvsAddress:         req.Info.AvsAddress,
			SlashAddr:          req.Info.SlashAddr,
			AvsOwnerAddress:    req.Info.AvsOwnerAddress,
			AssetId:            req.Info.AssetId,
			AvsUnbondingPeriod: req.Info.AvsUnbondingPeriod,
			MinSelfDelegation:  req.Info.MinSelfDelegation,
			OperatorAddress:    nil,
			EpochIdentifier:    req.Info.EpochIdentifier,
			StartingEpoch:      req.Info.StartingEpoch,
		},
	}

	res, err := sc.baseclient.SendTx(msg)

	return res.String(), err
}
