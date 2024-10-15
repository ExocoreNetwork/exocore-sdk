package main

import (
	"context"
	"encoding/hex"
	"fmt"
	sdkEcdsa "github.com/ExocoreNetwork/exocore-sdk/crypto/ecdsa"
	"log"
	"os"
	"path/filepath"

	key "github.com/ExocoreNetwork/exocore-sdk/cmd/exokey/generate"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	EthUrl       = "http://127.0.0.1:8545"
	EthWsUrl     = "ws://127.0.0.1:8546"
	KeystorePath = "/chainio/clients/tests/keys/test.ecdsa.key.json"
)

func main() {
	key.ImportKeyToFile("d196dca836f8ac2fff45b3c9f0113825ccbb33fa1b39737b948503b263ed75ae")
	ecdsaKeyPassword, _ := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")

	currentDir, err := os.Getwd()

	KeystorePath := filepath.Join(currentDir, KeystorePath)

	key, err := sdkEcdsa.ReadKey(KeystorePath, ecdsaKeyPassword)
	fmt.Println(hex.EncodeToString(key.D.Bytes()))
	client, err := ethclient.Dial(EthWsUrl)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xDF907c29719154eb9872f021d21CAE6E5025d7aB")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
