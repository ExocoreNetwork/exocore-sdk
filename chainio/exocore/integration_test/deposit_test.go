package integration_test

import (
	"fmt"
	deposit "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/x/deposit"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestDeposit() {
	cases := []SubTest{

		{
			"TestQueParams",
			QueryParams,
		},
		{
			"TestUpdateParams",
			UpdateParams,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func QueryParams(s IntegrationTestSuite) {

	res, err := s.Deposit.QueryParams(deposit.QueryParamsRequest{})
	require.NoError(s.T(), err)
	fmt.Println(res)

}
func UpdateParams(s IntegrationTestSuite) {

	a := deposit.Params{
		ExoCoreLzAppAddress: "0x561154bD8918c0bE7691F1BF8716AEd1CCC01d50", ExoCoreLzAppEventTopic: "0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec",
	}
	params := deposit.MsgUpdateParams{Authority: "", Params: a}

	res, err := s.Deposit.UpdateParams(params)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), res)

}
