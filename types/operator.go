package types

import (
	"encoding/json"
	"github.com/ExocoreNetwork/exocore-sdk/utils"
)

const (
	ZeroAddress = "0x0000000000000000000000000000000000000000"
)

// Operator represents EigenLayer's view of an operator
type Operator struct {
	// Address of the operator
	Address string `yaml:"address" json:"address"`

	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/delegation-redesign/src/contracts/interfaces/IDelegationManager.sol#L18
	EarningsReceiverAddress   string `yaml:"earnings_receiver_address"    json:"earnings_receiver_address"`
	DelegationApproverAddress string `yaml:"delegation_approver_address"  json:"delegation_approver_address"`
	StakerOptOutWindowBlocks  uint32 `yaml:"staker_opt_out_window_blocks" json:"staker_opt_out_window_blocks"`

	// MetadataUrl URL where operator metadata is stored
	MetadataUrl string `yaml:"metadata_url" json:"metadata_url"`
}

func (o Operator) Validate() error {
	if !utils.IsValidEthereumAddress(o.Address) {
		return ErrInvalidOperatorAddress
	}

	if !utils.IsValidEthereumAddress(o.EarningsReceiverAddress) {
		return ErrInvalidEarningsReceiverAddress
	}

	if o.DelegationApproverAddress != ZeroAddress && !utils.IsValidEthereumAddress(o.DelegationApproverAddress) {
		return ErrInvalidDelegationApproverAddress
	}

	err := utils.CheckIfUrlIsValid(o.MetadataUrl)
	if err != nil {
		return utils.WrapError(ErrInvalidMetadataUrl, err)
	}

	body, err := utils.ReadPublicURL(o.MetadataUrl)
	if err != nil {
		return utils.WrapError(ErrReadingMetadataUrlResponse, err)
	}

	operatorMetadata := OperatorMetadata{}
	err = json.Unmarshal(body, &operatorMetadata)
	if err != nil {
		return ErrUnmarshalOperatorMetadata
	}

	return operatorMetadata.Validate()
}
