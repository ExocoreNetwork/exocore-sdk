package crypto_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ExocoreNetwork/exocore-sdk/common/crypto"
	sdk "github.com/ExocoreNetwork/exocore-sdk/types"
)

func TestNewMnemonicKeyManager(t *testing.T) {
	mnemonic := "nerve leader thank marriage spice task van start piece crowd run hospital control outside cousin romance left choice poet wagon rude climb leisure spring"

	km, err := crypto.NewMnemonicKeyManager(mnemonic, "secp256k1")
	assert.NoError(t, err)

	pubKey := km.ExportPubKey()
	pubkeyBech32, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, pubKey)
	assert.NoError(t, err)
	assert.Equal(t, "iap1qf6rwt2vpsdx9tcwq03w4dw9udd657u0gmknjd4l0ht699x6npll6hf0ru9", pubkeyBech32)

	address := sdk.AccAddress(pubKey.Address()).String()
	assert.Equal(t, "iaa1y9kd9uy7a4qnjp0z5yjx5jhrkv2ycdkzqc0h8z", address)
}

func TestNewMnemonicKeyManager01(t *testing.T) {
	//mnemonic := "nerve leader thank marriage spice task van start piece crowd run hospital control outside cousin romance left choice poet wagon rude climb leisure spring"
	privateKey := "D35DA4AE8E66DD85C40E3C610B2D1F81CC6592FEF9E05A4A3308D05A6572022F"
	pkBytes, _ := hex.DecodeString(privateKey)
	km, err := crypto.NewPriKeyManager(pkBytes, "secp256k1")
	assert.NoError(t, err)

	pubKey := km.ExportPubKey()
	pubkeyBech32, err := sdk.Bech32ifyPubKey("evmospub", pubKey)
	assert.NoError(t, err)
	assert.Equal(t, "evmospub1qd4qud99e76q3xjzhr37rgx6a9juvkvuqprch6ph54hg3jywwfqnjrw3sd3", pubkeyBech32)

	address := sdk.AccAddress(pubKey.Address()).String()
	assert.Equal(t, "evmos1ya5zu5ayvghyznkwqqyer95v6efv8s5krs897u", address)
}
