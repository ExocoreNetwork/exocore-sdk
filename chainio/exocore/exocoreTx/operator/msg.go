package operator

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "deposit"
)

var (
	_ sdk.Msg = &RegisterOperatorReq{}
)

func (msg RegisterOperatorReq) Route() string { return ModuleName }

func (msg RegisterOperatorReq) Type() string { return "reg_operator" }

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (m *RegisterOperatorReq) GetSigners() []sdk.AccAddress {
	addr := sdk.MustAccAddressFromBech32(m.FromAddress)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check of the provided data
func (m *RegisterOperatorReq) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.FromAddress); err != nil {
		return errorsmod.Wrap(err, "invalid from address")
	}
	return nil
}

// GetSignBytes implements the LegacyMsg interface.
func (m *RegisterOperatorReq) GetSignBytes() []byte {
	return nil
}
