package common

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/evmos/evmos/v16/app"
	"github.com/evmos/evmos/v16/encoding"
	grpc1 "github.com/gogo/protobuf/grpc"
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
	Msgs          []sdk.Msg
	Sequence      uint64
	AccountNumber uint64
}

// SendTx delivers a cosmos tx for a given set of msgs
func SendTx(args CosmosTxArgs, grpcConn grpc1.ClientConn) (*txtypes.BroadcastTxResponse, error) {
	txConfig := encoding.MakeConfig(app.ModuleBasics).TxConfig
	tx, err := PrepareCosmosTx(
		CosmosTxArgs{
			TxCfg:         txConfig,
			Priv:          args.Priv,
			ChainID:       args.ChainID,
			Gas:           args.Gas,
			GasPrice:      args.GasPrice,
			Msgs:          args.Msgs,
			Sequence:      args.Sequence,
			AccountNumber: args.AccountNumber,
		},
	)
	if err != nil {
		return &txtypes.BroadcastTxResponse{}, err
	}
	return BroadcastTxBytes(txConfig.TxEncoder(), tx, grpcConn)
}

// BroadcastTxBytes encodes a transaction and calls SendTx on the app.
func BroadcastTxBytes(txEncoder sdk.TxEncoder, tx sdk.Tx, grpcConn grpc1.ClientConn) (*txtypes.BroadcastTxResponse, error) {
	// bz are bytes to be broadcasted over the network
	bz, err := txEncoder(tx)
	if err != nil {
		return &txtypes.BroadcastTxResponse{}, err
	}

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
		fees = sdk.Coins{sdk.NewCoin("aexo", sdk.NewIntFromUint64(uint64(1000000)))}
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
		Sequence: args.Sequence,
	}

	sigsV2 := []signing.SignatureV2{sigV2}

	if err := txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	signerData := authsigning.SignerData{
		ChainID:       args.ChainID,
		AccountNumber: args.AccountNumber,
		Sequence:      args.Sequence,
	}
	sigV2, err := tx.SignWithPrivKey(
		args.TxCfg.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder, args.Priv, args.TxCfg,
		args.Sequence,
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
