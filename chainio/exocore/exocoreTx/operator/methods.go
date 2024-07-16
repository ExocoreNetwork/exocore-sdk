package operator

import (
	"context"
	grpc "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/client"
	util "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type operatorClient struct {
	baseclient grpc.BaseClient
}

func (sc operatorClient) RegisterInterfaceTypes(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}
func (sc operatorClient) Name() string {
	return ModuleName
}

func NewClient(baseClient grpc.BaseClient) *operatorClient {
	return &operatorClient{
		baseclient: baseClient,
	}
}

func (sc operatorClient) GetOperatorInfo(req GetOperatorInfoReq) (OperatorInfo, error) {
	conn, err := sc.baseclient.GenConn()

	if err != nil {
		return OperatorInfo{}, err
	}

	res, err := NewQueryClient(conn).GetOperatorInfo(
		context.Background(),
		&req,
	)

	if err != nil {
		return OperatorInfo{}, err
	}
	return *res, nil
}

func (sc operatorClient) RegisterOperator(req RegisterOperatorReq) (string, error) {

	accountAddress := sdk.AccAddress(sc.baseclient.GenCfg().Priv.PubKey().Address().Bytes())
	msg := &RegisterOperatorReq{
		FromAddress: util.GetBech32Addr("exo", accountAddress),

		Info: &OperatorInfo{
			EarningsAddr:     "",
			ApproveAddr:      "",
			OperatorMetaInfo: "test operator",
			ClientChainEarningsAddr: &ClientChainEarningAddrList{
				EarningInfoList: nil,
			},
		},
	}

	res, err := sc.baseclient.SendTx(msg)

	return res.String(), err
}
