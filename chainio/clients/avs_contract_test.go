package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	sdkexocontracts "github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exocontracts"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	avssub "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prysmaticlabs/prysm/v4/crypto/bls/blst"
	blscommon "github.com/prysmaticlabs/prysm/v4/crypto/bls/common"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

type ExoClientService struct {
	ethClient     eth.EthClient
	exocoreReader sdkexocontracts.EXOReader
	exocoreWriter sdkexocontracts.EXOWriter
	exocoreSub    sdkexocontracts.AvsRegistryChainSubscriber
}

const (
	EthUrl       = "http://127.0.0.1:8545"
	EthWsUrl     = "ws://127.0.0.1:8546"
	KeystorePath = "/tests/keys/test.ecdsa.key.json"
	AvsAddress   = "0x52ce3752aF2A5675F23DA541e6cB78581369362E"
)

func NewExoClientService() (*ExoClientService, error) {

	//privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")

	var ethRpcClient eth.EthClient
	ethRpcClient, err := eth.NewClient(EthUrl)
	if err != nil {
		fmt.Println("Cannot create http ethclient", "err", err)
		return nil, err
	}
	ethWsClient, err := eth.NewClient(EthWsUrl)
	if err != nil {
		fmt.Println("Cannot create ws ethclient", "err", err)
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

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	KeystorePath := filepath.Join(currentDir, KeystorePath)
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

	exoChainWriter, _ := sdkexocontracts.BuildELChainWriter(
		common.HexToAddress(AvsAddress),

		ethRpcClient,
		logger,
		txMgr)

	exoChainSub, _ := sdkexocontracts.BuildAvsRegistryChainSubscriber(
		common.HexToAddress(AvsAddress),
		ethWsClient,
		logger,
	)
	exoClientService := &ExoClientService{
		ethClient:     ethRpcClient,
		exocoreReader: exoChainReader,
		exocoreWriter: exoChainWriter,
		exocoreSub:    *exoChainSub,
	}
	return exoClientService, nil
}

// sender addr is 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5(exo1sc9kjykz6qehauzmhjympsktdjaw4d99dksgrk)
func TestRunRegisterAVS(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	avsName, epochIdentifier := "avsTest", "day"
	avsOwnerAddress := []string{"exo13h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "exo1sc9kjykz6qehauzmhjympsktdjaw4d99dksgrk"}
	assetIds := []string{"0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"}
	minStakeAmount, avsUnbondingPeriod, minSelfDelegation := 3, 3, 5
	params := []uint64{5, 7, 8, 4}
	_, err := service.exocoreWriter.RegisterAVSToExocore(context.Background(),
		avsName,
		uint64(minStakeAmount),
		gethcommon.HexToAddress("0xDF907c29719154eb9872f021d21CAE6E5025d7aB"),
		gethcommon.HexToAddress("0xDF907c29719154eb9872f021d21CAE6E5025d7aB"),
		gethcommon.HexToAddress("0xDF907c29719154eb9872f021d21CAE6E5025d7aB"),
		avsOwnerAddress,
		assetIds,
		uint64(avsUnbondingPeriod),
		uint64(minSelfDelegation),
		epochIdentifier,
		params,
	)
	fmt.Println(err)
}

func TestRunSubTask(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	newTaskCreatedChan := make(chan *avssub.ContractavsserviceTaskCreated)
	sub := service.exocoreSub.SubscribeToNewTasks(newTaskCreatedChan)
	i := 1
	for {
		fmt.Println(i)
		i++
		select {
		case <-context.Background().Done():
			fmt.Println("done")
		case err := <-sub.Err():
			fmt.Println("Error in websocket subscription", "err", err)
			sub.Unsubscribe()
			sub = service.exocoreSub.SubscribeToNewTasks(newTaskCreatedChan)
		case newTaskCreatedLog := <-newTaskCreatedChan:
			fmt.Println("newTaskCreatedChan：", newTaskCreatedLog)
		}
	}
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

func TestGetBalance(t *testing.T) {

	client, err := ethclient.Dial(EthUrl)
	if err != nil {
		panic(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("Cannot get chainId", "err", err)
		return
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		fmt.Println("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	KeystorePath := filepath.Join(currentDir, KeystorePath)
	_, senderAddress, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: KeystorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	balance, err := client.BalanceAt(context.Background(), senderAddress, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(senderAddress.String())
	fmt.Printf("balance = %v\n", balance)

}

func Test_eth_subscribe(t *testing.T) {

	client, err := eth.NewClient(EthWsUrl)
	if err != nil {
		fmt.Println("Cannot create ws ethclient", "err", err)
	}

	contractAddress := common.HexToAddress(AvsAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
func Test_bls(t *testing.T) {
	msg := crypto.Keccak256Hash([]byte("7"))
	privateKeys := make([]blscommon.SecretKey, 4)
	publicKeys := make([]blscommon.PublicKey, 4)

	sigs := make([]blscommon.Signature, 4)
	for i := 0; i < 4; i++ {
		privateKey, _ := blst.RandKey()
		pri := privateKey.Marshal()
		pub := privateKey.PublicKey().Marshal()
		fmt.Println("pri:", hex.EncodeToString(pri))
		fmt.Println("pub:", hex.EncodeToString(pub))

		privateKeys[i] = privateKey
		publicKeys[i] = privateKey.PublicKey()
		fmt.Println(len(privateKey.PublicKey().Marshal()))
		sigs[i] = privateKey.Sign(msg[:])
		dsig := privateKey.Sign(msg[:]).Marshal()
		fmt.Println(len(dsig))

		fmt.Println("dsig:", hex.EncodeToString(dsig))
		valid := privateKey.Sign(msg[:]).Verify(privateKey.PublicKey(), msg[:])
		require.True(t, valid, "Signature verification failed")

	}
	//s.publicKeys = append(s.publicKeys[:1], s.publicKeys[1+1:]...)

	aggregatedPublicKey := blst.AggregateMultiplePubkeys(publicKeys)
	aggregatedPub := aggregatedPublicKey.Marshal()
	fmt.Println("aggregatedPublicKey:", hex.EncodeToString(aggregatedPub))

	signature := blst.AggregateSignatures(sigs)

	valid := signature.FastAggregateVerify(publicKeys, msg)
	valid1 := signature.Verify(aggregatedPublicKey, msg.Bytes())

	require.True(t, valid, "Signature verification failed")
	require.True(t, valid1, "Signature verification failed")

	fmt.Println("msg:", fmt.Sprintf("%x", msg)) //ok
	aggsig := signature.Marshal()
	fmt.Println("aggsignature:", hex.EncodeToString(aggsig)) //ok

}
func Test_bls_msgs(t *testing.T) {

	privateKeys := make([]blscommon.SecretKey, 3)
	for i := 0; i < 3; i++ {
		privateKeys[i], _ = blst.RandKey()
	}

	publicKeys := make([]blscommon.PublicKey, 3)
	for i := 0; i < 3; i++ {
		publicKeys[i] = privateKeys[i].PublicKey()
	}

	messages := [][32]byte{
		crypto.Keccak256Hash([]byte("1")),
		crypto.Keccak256Hash([]byte("2")),
		crypto.Keccak256Hash([]byte("3")),
	}

	signatures := make([]blscommon.Signature, 3)
	for i := 0; i < 3; i++ {
		msg := messages[i][:]
		signatures[i] = privateKeys[i].Sign(msg)
	}

	aggsignature := blst.AggregateSignatures(signatures)

	valid1 := aggsignature.AggregateVerify(publicKeys, messages)
	valid2 := aggsignature.FastAggregateVerify(publicKeys, crypto.Keccak256Hash([]byte("1")))
	valid3 := aggsignature.Eth2FastAggregateVerify(publicKeys, crypto.Keccak256Hash([]byte("1")))
	fmt.Println("Aggregate signature1 is valid for all messages:", valid1)
	fmt.Println("Aggregate signature2 is valid for all messages:", valid2)
	fmt.Println("Aggregate signature3 is valid for all messages:", valid3)

	require.True(t, valid1, "Signature verification failed")

}

func Test_bls_msgs1(t *testing.T) {

	privateKeys := make([]blscommon.SecretKey, 3)
	for i := 0; i < 3; i++ {
		privateKeys[i], _ = blst.RandKey()
	}

	publicKeys := make([]blscommon.PublicKey, 3)
	for i := 0; i < 3; i++ {
		publicKeys[i] = privateKeys[i].PublicKey()
	}
	msg := crypto.Keccak256Hash([]byte("1"))

	signatures := make([]blscommon.Signature, 3)
	for i := 0; i < 3; i++ {
		signatures[i] = privateKeys[i].Sign(msg.Bytes())
	}

	aggsignature := blst.AggregateSignatures(signatures)

	valid2 := aggsignature.FastAggregateVerify(publicKeys, msg)
	valid3 := aggsignature.Eth2FastAggregateVerify(publicKeys, msg)
	fmt.Println("Aggregate signature2 is valid for all messages:", valid2)
	fmt.Println("Aggregate signature3 is valid for all messages:", valid3)
}
