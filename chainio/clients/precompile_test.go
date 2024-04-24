package main

import (
	"context"
	"fmt"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	sdkexocontracts "github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exoprecompile"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"testing"
)

type ExoClientService struct {
	ethClient     eth.EthClient
	exocoreReader sdkexocontracts.EXOReader
	exocoreWriter sdkexocontracts.EXOWriter
}

const (
	EthUrl                   = "http://38.91.106.205:8545"
	KeystorePath             = "/Users/devin/2023/blockchain/EXO/exocore-sdk/chainio/clients/tests/keys/test.ecdsa.key.json"
	DepositPrecompileAddress = "0xf84e8EfCd48904D374652B637c63762aeE4Ea6ab"
	AvsPrecompileAddress     = "0x970b60FbF35F2EFCFdf82Ee870ab9F267B238883"
	AvsTaskPrecompileAddress = "0xf84e8EfCd48904D374652B637c63762aeE4Ea611"
	BlsPrecompileAddress     = "0xf84e8EfCd48904D374652B637c63762aeE4Ea622"
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

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, nil, signerV2, senderAddress)

	// get the Reader for the Exo contracts
	exoChainReader, _ := sdkexocontracts.BuildExoChainReader(
		common.HexToAddress(DepositPrecompileAddress),
		common.HexToAddress(AvsPrecompileAddress),
		common.HexToAddress(AvsTaskPrecompileAddress),
		common.HexToAddress(BlsPrecompileAddress),
		ethRpcClient,
		nil)

	elChainWriter, _ := sdkexocontracts.BuildELChainWriter(
		common.HexToAddress(DepositPrecompileAddress),
		common.HexToAddress(AvsPrecompileAddress),
		common.HexToAddress(AvsTaskPrecompileAddress),
		common.HexToAddress(BlsPrecompileAddress),
		ethRpcClient,
		nil,
		txMgr)

	exoClientService := &ExoClientService{
		ethClient:     ethRpcClient,
		exocoreReader: exoChainReader,
		exocoreWriter: elChainWriter,
	}
	return exoClientService, nil
}
func TestRunDepositTo(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	clientChainLzID := 101
	usdtAddress := utils.PaddingClientChainAddress(common.FromHex("0xdAC17F958D2ee523a2206206994597C13D831ec7"), 32)
	stakerAddr := utils.PaddingClientChainAddress(common.FromHex("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"), 32)
	opAmount := big.NewInt(100)

	_, err := service.exocoreWriter.DepositTo(context.Background(), uint16(clientChainLzID), usdtAddress, stakerAddr, opAmount)
	fmt.Println(err)
}

func TestRunAVSAction(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	avsName, avsAddress, operatorAddress, avsOwnerAddress, assetID := "avsTest", "exo13h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "exo18h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "0xdAC17F958D2ee523a2206206994597C13D831ec7", ""
	_, err := service.exocoreWriter.AVSAction(context.Background(), avsName, avsAddress, operatorAddress, 1, avsOwnerAddress, assetID)
	fmt.Println(err)
}

func TestGetCode(t *testing.T) {

	client, err := ethclient.Dial(EthUrl)
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress(DepositPrecompileAddress)

	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Contract code = %v\n", code)

}
