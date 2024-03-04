package integration_test

import (
	sdk "github.com/ExocoreNetwork/exocore-sdk/chainio/exocore"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/exocore/types"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

const (
	nodeURI    = "http://23.162.56.84:1317"
	grpcAddr   = "23.162.56.84:9090"
	chainID    = "exocoretestnet_233-1"
	privateKey = "7A63E68134660B0098A1BDC2A22E36DC6B5B428507B7D2A684F44D3CC3290616"
	gasPrice   = 9
	gas        = 1000000
)

type IntegrationTestSuite struct {
	suite.Suite
	sdk.Client
	r *rand.Rand
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	bech32AddressPrefix := types.AddrPrefixCfg{
		AccountAddr:   "exo",
		ValidatorAddr: "exovaloper",
		ConsensusAddr: "exovalcons",
		AccountPub:    "exopub",
		ValidatorPub:  "exovaloperpub",
		ConsensusPub:  "exovalconspub",
	}
	options := []types.Option{
		types.TimeoutOption(10),
		types.Bech32AddressPrefixOption(&bech32AddressPrefix),
		types.BIP44PathOption(""),
	}
	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, privateKey, gas, gasPrice, options...)
	if err != nil {
		panic(err)
	}

	s.Client = sdk.NewClient(cfg)
	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
