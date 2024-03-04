package deposit

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/evmos/evmos/v14/app"
	"github.com/evmos/evmos/v14/crypto/ethsecp256k1"
	"github.com/evmos/evmos/v14/encoding"
	grpc1 "github.com/gogo/protobuf/grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"math"
	"strings"
	"testing"
)

var (
	feeAmt     = math.Pow10(16)
	DefaultFee = sdk.NewCoin("aexo", sdk.NewIntFromUint64(uint64(1000000))) // 0.01 aexo
)

const (
	/*	memo          = "testmemo"
		gasLimit      = uint64(200000)
		timeoutHeight = uint64(9999999999)*/
	URL1       = "23.162.56.84:9090"
	URL        = "localhost:9090"
	privateKey = "62E8DCAE790AB9BE9EB3A58950791F9A989BF4649F980EC6BF7700B0518BAEAC"
	ChainID    = "exocoretestnet_233-1"
	seq        = uint64(4)
	accNumber  = uint64(0)
)

// CosmosTxArgs contains the params to create a cosmos tx
type CosmosTxArgs struct {
	// TxCfg is the client transaction config
	TxCfg client.TxConfig
	// Priv is the private key that will be used to sign the tx
	Priv cryptotypes.PrivKey
	// ChainID is the chain's id on cosmos format, e.g. 'evmos_9000-1'
	ChainID string
	// Gas to be used on the tx
	Gas uint64
	// GasPrice to use on tx
	GasPrice *sdkmath.Int
	// Fees is the fee to be used on the tx (amount and denom)
	Fees sdk.Coins
	// FeeGranter is the account address of the fee granter
	FeeGranter sdk.AccAddress
	//
	Memo string
	// Msgs slice of messages to include on the tx
	Msgs []sdk.Msg
}

func TestExoDelegate(t *testing.T) {
	keystr, _ := hex.DecodeString(privateKey)
	key := ethsecp256k1.PrivKey{Key: keystr}

	conn := GetGRPCClientforexo(URL)
	queryClient := stakingtypes.NewQueryClient(conn)
	validator, _ := queryClient.Validator(context.Background(), &stakingtypes.QueryValidatorRequest{ValidatorAddr: "exovaloper1xpcan9zdtlsku47fpxdjlkvm669c2dftdrvg30"})

	_, err := Delegate(&key, DefaultFee, validator.Validator)
	require.NoError(t, err)
}

func TestExoTx(t *testing.T) {
	keystr, _ := hex.DecodeString(privateKey)
	key := ethsecp256k1.PrivKey{Key: keystr}

	res, _ := QueryParams(QueryParamsRequest{})
	fmt.Println(res)
	UpdateParams(&key)
}

func QueryParams(req QueryParamsRequest) (QueryParamsResponse, error) {
	conn := GetGRPCClientforexo(URL)

	res, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)

	if err != nil {
		return QueryParamsResponse{}, err
	}
	return QueryParamsResponse{res.Params}, nil
}

func UpdateParams(priv *ethsecp256k1.PrivKey) (*txtypes.BroadcastTxResponse, error) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())
	msg := &MsgUpdateParams{
		Authority: getBech32Addr("exo", accountAddress),
		Params: Params{
			ExoCoreLzAppAddress:    "0x561154bD8918c0bE7691F1BF8716AEd1CCC01de4",
			ExoCoreLzAppEventTopic: "0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec",
		},
	}

	intVal := sdk.NewInt(9)
	intPtr := &intVal

	return DeliverTx(priv, intPtr, []sdk.Msg{msg}...)
}

// Delegate delivers a delegate tx
func Delegate(
	priv *ethsecp256k1.PrivKey,
	delegateAmount sdk.Coin,
	validator stakingtypes.Validator,
) (*txtypes.BroadcastTxResponse, error) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())

	val, err := ValAddressFromBech32(validator.OperatorAddress)
	if err != nil {
		return &txtypes.BroadcastTxResponse{}, err
	}

	delegateMsg := stakingtypes.NewMsgDelegate(accountAddress, sdk.ValAddress(val), delegateAmount)
	delegateMsg.DelegatorAddress = getBech32Addr("exo", accountAddress)
	delegateMsg.ValidatorAddress = getBech32Addr("exo", val)
	intVal := sdk.NewInt(2)
	intPtr := &intVal

	return DeliverTx(priv, intPtr, delegateMsg)
}

// DeliverTx delivers a cosmos tx for a given set of msgs
func DeliverTx(priv cryptotypes.PrivKey, gasPrice *sdkmath.Int, msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	txConfig := encoding.MakeConfig(app.ModuleBasics).TxConfig
	tx, err := PrepareCosmosTx(
		CosmosTxArgs{
			TxCfg:    txConfig,
			Priv:     priv,
			ChainID:  ChainID,
			Gas:      1000000,
			GasPrice: gasPrice,
			Msgs:     msgs,
		},
	)
	if err != nil {
		return &txtypes.BroadcastTxResponse{}, err
	}
	return BroadcastTxBytes(txConfig.TxEncoder(), tx)
}

// BroadcastTxBytes encodes a transaction and calls DeliverTx on the app.
func BroadcastTxBytes(txEncoder sdk.TxEncoder, tx sdk.Tx) (*txtypes.BroadcastTxResponse, error) {
	// bz are bytes to be broadcasted over the network
	bz, err := txEncoder(tx)
	if err != nil {
		return &txtypes.BroadcastTxResponse{}, err
	}

	grpcConn := GetGRPCClientforexo(URL)
	txClient := txtypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.

	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: bz, // Proto-binary of the signed transaction, see previous step.
		},
	)

	return grpcRes, nil
}

// PrepareCosmosTx creates a cosmos tx and signs it with the provided messages and private key.
// It returns the signed transaction and an error
func PrepareCosmosTx(args CosmosTxArgs) (authsigning.Tx, error) {
	txBuilder := args.TxCfg.NewTxBuilder()

	txBuilder.SetGasLimit(args.Gas)

	var fees sdk.Coins
	if args.GasPrice != nil {
		fees = sdk.Coins{{Denom: "aexo", Amount: args.GasPrice.MulRaw(int64(args.Gas))}}
	} else {
		fees = sdk.Coins{DefaultFee}
	}

	txBuilder.SetFeeAmount(fees)
	if err := txBuilder.SetMsgs(args.Msgs...); err != nil {
		return nil, err
	}

	txBuilder.SetFeeGranter(args.FeeGranter)
	return signCosmosTx(
		args,
		txBuilder,
	)
}
func signCosmosTx(args CosmosTxArgs, txBuilder client.TxBuilder) (authsigning.Tx, error) {
	// addr := sdk.AccAddress(args.Priv.PubKey().Address().Bytes())
	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sigV2 := signing.SignatureV2{
		PubKey: args.Priv.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  args.TxCfg.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: seq,
	}

	sigsV2 := []signing.SignatureV2{sigV2}

	if err := txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	signerData := authsigning.SignerData{
		ChainID:       args.ChainID,
		AccountNumber: accNumber,
		Sequence:      seq,
	}
	sigV2, err := tx.SignWithPrivKey(
		args.TxCfg.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder, args.Priv, args.TxCfg,
		seq,
	)
	if err != nil {
		return nil, err
	}

	sigsV2 = []signing.SignatureV2{sigV2}
	if err = txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}
	return txBuilder.GetTx(), nil
}
func GetGRPCClientforexo(url string, options ...grpc.DialOption) grpc1.ClientConn {
	if options == nil {
		options = []grpc.DialOption{grpc.WithInsecure()}
	}

	clientConn, err := grpc.Dial(url, options...)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return clientConn
}

type ValAddress []byte

// ValAddressFromBech32 creates a ValAddress from a Bech32 string.
func ValAddressFromBech32(address string) (addr ValAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return ValAddress{}, errors.New("empty address string is not allowed")
	}

	bech32PrefixValAddr := "exovaloper"

	bz, err := GetFromBech32(address, bech32PrefixValAddr)
	if err != nil {
		return nil, err
	}

	return ValAddress(bz), nil
}

// GetFromBech32 decodes a bytestring from a Bech32 encoded string.
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, nil
	}

	hrp, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid Bech32 prefix; expected %s, got %s", prefix, hrp)
	}

	return bz, nil
}
func getBech32Addr(prefix string, addr []byte) string {
	bech32Addr, err := bech32.ConvertAndEncode(prefix, addr)
	if err != nil {
		panic(err)
	}
	return bech32Addr
}
