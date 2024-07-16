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
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	EthUrl       = "http://api-eth.exocore-restaking.com"
	EthWsUrl     = "ws://api-eth-wss.exocore-restaking.com"
	KeystorePath = "/tests/keys/test.ecdsa.key.json"
	AvsAddress   = "0x1cC54237A0b0804Af90D8078E9a224a0EB636999"
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
	operator, err := service.exocoreReader.Operators(nil, gethcommon.HexToAddress("0x2766fb07398cdfc4b0043de68dd8a8715bcbd955"))

	fmt.Println(operator)
	fmt.Println(err)
}

func TestRunSubTask(t *testing.T) {
	service, _ := NewExoClientService()
	fmt.Println(service.ethClient.ChainID(context.Background()))
	newTaskCreatedChan := make(chan *avssub.ContractavsserviceTaskCreated)
	operator, _ := service.exocoreReader.Operators(nil, gethcommon.HexToAddress("0x2766fb07398cdfc4b0043de68dd8a8715bcbd955"))
	fmt.Println(operator)
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
		sigs[i] = privateKey.Sign(msg[:])
		dsig := privateKey.Sign(msg[:]).Marshal()
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
	require.True(t, valid, "Signature verification failed")

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

	valid := aggsignature.AggregateVerify(publicKeys, messages)
	require.True(t, valid, "Signature verification failed")

	fmt.Println("Aggregate signature is valid for all messages:", valid)
}
