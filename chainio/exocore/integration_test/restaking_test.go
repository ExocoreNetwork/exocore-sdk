package integration_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
)

func (s IntegrationTestSuite) TestReStaking() {
	cases := []SubTest{
		{
			"TestQueClientChainInfoByIndex",
			QueClientChainInfoByIndex,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func QueClientChainInfoByIndex(s IntegrationTestSuite) {
	res, err := s.AssetManageClient.QueClientChainInfoByIndex(101)
	require.NoError(s.T(), err)
	fmt.Println(res)
}
