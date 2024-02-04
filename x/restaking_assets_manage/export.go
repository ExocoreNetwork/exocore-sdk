package restaking_assets_manage

import (
	sdk "github.com/ExocoreNetwork/exocore-sdk/types"
	"time"
)

// expose Gov module api for user
type Client interface {
	sdk.Module
	SubmitProposal(request SubmitProposalRequest, baseTx sdk.BaseTx) (uint64, sdk.ResultTx, sdk.Error)
	QueryProposal(proposalId uint64) (QueryProposalResp, sdk.Error)
	QueryParams(paramsType string) (QueryParamsResp, sdk.Error)
}

type SubmitProposalRequest struct {
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Type           string       `json:"type"`
	InitialDeposit sdk.DecCoins `json:"initial_deposit"`
}

type QueryProposalResp struct {
	ProposalId      uint64    `json:"proposal_id"`
	Status          string    `json:"status"`
	SubmitTime      time.Time `json:"submit_time"`
	DepositEndTime  time.Time `json:"deposit_end_time"`
	TotalDeposit    sdk.Coins `json:"total_deposit"`
	VotingStartTime time.Time `json:"voting_start_time"`
	VotingEndTime   time.Time `json:"voting_end_time"`
}

type (
	votingParams struct {
		VotingPeriod time.Duration `json:"voting_period"`
	}
	depositParams struct {
		MinDeposit       sdk.Coins     `json:"min_deposit"`
		MaxDepositPeriod time.Duration `json:"max_deposit_period"`
	}
	tallyParams struct {
		Quorum        sdk.Dec `json:"quorum"`
		Threshold     sdk.Dec `json:"threshold"`
		VetoThreshold sdk.Dec `json:"veto_threshold"`
	}
	QueryParamsResp struct {
		VotingParams  votingParams  `json:"voting_params"`
		DepositParams depositParams `json:"deposit_params"`
		TallyParams   tallyParams   `json:"tally_params"`
	}
)
