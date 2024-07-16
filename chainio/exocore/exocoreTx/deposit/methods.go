package deposit

import (
	"context"
	grpc "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/client"
	util "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type depositClient struct {
	baseclient grpc.BaseClient
}

func (sc depositClient) RegisterInterfaceTypes(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}
func (sc depositClient) Name() string {
	return ModuleName
}

func NewClient(baseClient grpc.BaseClient) *depositClient {
	return &depositClient{
		baseclient: baseClient,
	}
}

func (sc depositClient) QueryParams(req QueryParamsRequest) (QueryParamsResponse, error) {
	conn, err := sc.baseclient.GenConn()

	if err != nil {
		return QueryParamsResponse{}, err
	}

	res, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)

	if err != nil {
		return QueryParamsResponse{}, err
	}
	return QueryParamsResponse{res.Params}, nil
}

func (sc depositClient) UpdateParams(req MsgUpdateParams) (string, error) {

	accountAddress := sdk.AccAddress(sc.baseclient.GenCfg().Priv.PubKey().Address().Bytes())
	msg := &MsgUpdateParams{
		Authority: util.GetBech32Addr("exo", accountAddress),
		Params: Params{
			ExoCoreLzAppAddress:    req.Params.ExoCoreLzAppAddress,
			ExoCoreLzAppEventTopic: req.Params.ExoCoreLzAppEventTopic,
		},
	}

	res, err := sc.baseclient.SendTx(msg)

	return res.String(), err
}
