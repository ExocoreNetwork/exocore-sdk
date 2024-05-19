package main

import (
	"context"
	"fmt"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	sdkexocontracts "github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exocontracts"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"testing"
)

type ExoClientService struct {
	ethClient     eth.EthClient
	exocoreReader sdkexocontracts.EXOReader
	exocoreWriter sdkexocontracts.EXOWriter
}

const (
	EthUrl       = "http://127.0.0.1:8545"
	KeystorePath = "/Users/trestin/go-latest/Project/2024/exocore-sdk-dev/chainio/clients/tests/keys/test.ecdsa.key.json"
	AvsAddress   = "0xDF907c29719154eb9872f021d21CAE6E5025d7aB"
)

func NewExoClientService() (*ExoClientService, error) {

	//privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")

	var ethRpcClient eth.EthClient
	ethRpcClient, err := eth.NewClient(EthUrl)
	if err != nil {
		fmt.Println("Cannot create http ethclient", "err", err)
		return nil, err
	}
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		fmt.Println("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		fmt.Println("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, senderAddress, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: KeystorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	logger, err := logging.NewZapLogger(logging.Development)

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, senderAddress)

	// get the Reader for the Exo contracts
	exoChainReader, _ := sdkexocontracts.BuildExoChainReader(
		common.HexToAddress(AvsAddress),
		ethRpcClient,
		logger)

	elChainWriter, _ := sdkexocontracts.BuildELChainWriter(
		common.HexToAddress(AvsAddress),

		ethRpcClient,
		logger,
		txMgr)

	exoClientService := &ExoClientService{
		ethClient:     ethRpcClient,
		exocoreReader: exoChainReader,
		exocoreWriter: elChainWriter,
	}
	return exoClientService, nil
}

func TestRunAVSAction(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	avsName, avsAddress, operatorAddress, avsOwnerAddress, assetID := "avsTest", "exo13h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "exo18h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "0xdAC17F958D2ee523a2206206994597C13D831ec7", "222"
	_, err := service.exocoreWriter.RegisterAVSToExocore(context.Background(), avsName, avsAddress, operatorAddress, 1, avsOwnerAddress, assetID)
	fmt.Println(err)
}
func TestRunRegisterOperatorToAVS(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	pubkey, err := hexutil.Decode("0x818c91204445f4052ae7f51f9fb85d67359b309cf98bc6d0e36f8b2916b50512666cfd2f74b1d2320e4648f4acadd282")
	r, err := service.exocoreWriter.RegisterOperatorToAVS(context.Background(), pubkey, "operator02")

	fmt.Println(r.Status)
	fmt.Println(err)
}

func TestRungetoperatorbyaddress(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	operator, err := service.exocoreReader.Operators(nil, gethcommon.HexToAddress("0x860b6912c2d0337ef05bbc89b0c2cb6cbaeab4a5"))

	fmt.Println(operator)
	fmt.Println(err)
}

func TestGetCode(t *testing.T) {

	client, err := ethclient.Dial(EthUrl)
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress(AvsAddress)

	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Contract code = %v\n", code)

}
