package client

import (
	"encoding/hex"
	"encoding/json"
	comm "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/common"
	sdktypes "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"
	util "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/ethereum/go-ethereum/log"
	"github.com/evmos/evmos/v16/crypto/ethsecp256k1"
	rest "github.com/go-resty/resty/v2"
	grpc1 "github.com/gogo/protobuf/grpc"
	"strconv"
)

type BaseClient struct {
	cfg            *sdktypes.ClientConfig
	encodingConfig sdktypes.EncodingConfig
	clientConn     GrpcClient
	restClient     RestClient
}

// NewBaseClient return the baseClient for every sub modules
func NewBaseClient(cfg sdktypes.ClientConfig, encodingConfig sdktypes.EncodingConfig, logger log.Logger) BaseClient {
	keystr, _ := hex.DecodeString(cfg.KeyStr)
	pri := ethsecp256k1.PrivKey{Key: keystr}
	cfg.Priv = &pri
	base := BaseClient{
		cfg:            &cfg,
		encodingConfig: encodingConfig,
		clientConn:     NewGRPCClient(cfg.GRPCAddr, cfg.GRPCOptions...),
		restClient:     NewRestClient(),
	}
	return base
}

func (base *BaseClient) GenConn() (grpc1.ClientConn, error) {
	return base.clientConn.GenConn()
}
func (base *BaseClient) GenRestConn() (rest.Client, error) {
	return base.restClient.GenRestConn()
}

func (base *BaseClient) GenCfg() *sdktypes.ClientConfig {
	return base.cfg
}

// SendTx delivers a cosmos tx for a given set of msgs
func (base *BaseClient) SendTx(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	intVal := sdk.NewIntFromUint64(base.GenCfg().GasPrice)
	grpcConn, _ := base.GenConn()
	restConn, _ := base.GenRestConn()

	accountAddress := sdk.AccAddress(base.GenCfg().Priv.PubKey().Address().Bytes())

	url := base.cfg.RPCAddr + "/cosmos/auth/v1beta1/accounts/" + util.GetBech32Addr("exo", accountAddress)
	resp, _ := restConn.R().Get(url)
	var response EthAccountResponse
	json.Unmarshal(resp.Body(), &response)
	seq, _ := strconv.ParseUint(response.Account.BaseAcc.Sequence, 10, 64)
	num, _ := strconv.ParseUint(response.Account.BaseAcc.AccountNumber, 10, 64)

	args := comm.CosmosTxArgs{
		Priv:          base.GenCfg().Priv,
		ChainID:       base.GenCfg().ChainID,
		Gas:           base.GenCfg().Gas,
		GasPrice:      &intVal,
		Msgs:          msgs,
		Sequence:      seq,
		AccountNumber: num,
	}

	res, _ := comm.SendTx(args, grpcConn)
	return res, nil
}

type EthAccountResponse struct {
	Account  EthAccount `json:"account"`
	CodeHash string     `json:"code_hash"`
}

type EthAccount struct {
	Type    string      `json:"@type"`
	BaseAcc BaseAccount `json:"base_account"`
}

type BaseAccount struct {
	Address       string `json:"address"`
	PubKey        PubKey `json:"pub_key"`
	AccountNumber string `json:"account_number"`
	Sequence      string `json:"sequence"`
}

type PubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}
