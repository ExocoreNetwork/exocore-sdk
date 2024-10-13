package clients_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	sdkexocontracts "github.com/ExocoreNetwork/exocore-sdk/chainio/clients/exocontracts"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	avssub "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/avs"
	sdkEcdsa "github.com/ExocoreNetwork/exocore-sdk/crypto/ecdsa"
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
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
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
)

func DeployAVS() (common.Address, string, error) {
	var ethRpcClient eth.EthClient
	ethRpcClient, err := eth.NewClient(EthUrl)
	if err != nil {
		log.Println("Cannot create http ethclient", "err", err)
		return common.Address{}, "", err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		log.Println("Cannot get chainId", "err", err)
		return common.Address{}, "", err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Println("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return common.Address{}, "", err
	}

	KeystorePath := filepath.Join(currentDir, KeystorePath)

	if err != nil {
		panic(err)
	}

	logger, err := logging.NewZapLogger(logging.Development)

	// addresses := []gethcommon.Address{
	//	common.HexToAddress("0x0000000000000000000000000000000000000000"),
	//	common.HexToAddress("0x1111111111111111111111111111111111111111"),
	//	common.HexToAddress("0x2222222222222222222222222222222222222222"),
	// }
	key, err := sdkEcdsa.ReadKey(KeystorePath, ecdsaKeyPassword)
	avsAddr, tx, err := sdkexocontracts.DeployAVS(
		ethRpcClient,
		logger,
		*key,
		chainId,
	)
	if err != nil {
		panic(err)
	}
	// get the Reader for the Exo contracts

	return avsAddr, tx, nil
}
func NewExoClientService(avsAddr common.Address) (*ExoClientService, error) {

	// privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")

	var ethRpcClient eth.EthClient
	ethRpcClient, err := eth.NewClient(EthUrl)
	if err != nil {
		log.Println("Cannot create http ethclient", "err", err)
		return nil, err
	}
	ethWsClient, err := eth.NewClient(EthWsUrl)
	if err != nil {
		log.Println("Cannot create ws ethclient", "err", err)
		return nil, err
	}
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		log.Println("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Println("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
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

	if err != nil {
		panic(err)
	}
	// get the Reader for the Exo contracts

	exoChainReader, _ := sdkexocontracts.BuildExoChainReader(
		avsAddr,
		ethRpcClient,
		logger)

	exoChainWriter, _ := sdkexocontracts.BuildELChainWriter(
		avsAddr,
		ethRpcClient,
		logger,
		txMgr)

	exoChainSub, _ := sdkexocontracts.BuildAvsRegistryChainSubscriber(
		avsAddr,
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

func generateRandomAddress() common.Address {
	// Generate a random private key
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	// Get the public key from the private key
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key to ECDSA")
	}

	// Convert the public key to an Ethereum address
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address
}

// sender addr is 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5(exo1sc9kjykz6qehauzmhjympsktdjaw4d99dksgrk)
// query result :exocored query avs AVSInfo 0x5C544D806EA31cfc1d9B6d458a72ed0b8a60Dd54
func TestRunRegisterAVS(t *testing.T) {
	avsAddr, _, err := DeployAVS()
	// The successful execution of the transaction for deploying the contract
	// and packaging it into the block requires waiting for a moment
	time.Sleep(5 * time.Second)
	log.Println("avs###" + avsAddr.String())
	service, _ := NewExoClientService(avsAddr)
	log.Println(service.ethClient.ChainID(context.Background()))
	avsName, epochIdentifier := "avsTest", "day"
	avsOwnerAddress := []string{"exo13h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "exo1sc9kjykz6qehauzmhjympsktdjaw4d99dksgrk"}
	assetIds := []string{"0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"}
	minStakeAmount, avsUnbondingPeriod, minSelfDelegation := 3, 3, 5
	params := []uint64{5, 7, 8, 4}
	_, err = service.exocoreWriter.RegisterAVSToExocore(context.Background(),
		avsName,
		uint64(minStakeAmount),
		generateRandomAddress(),
		gethcommon.HexToAddress("0x92D203486fc326Eaad87f3B876b3A5a7db245F3c"),
		gethcommon.HexToAddress("0x92D203486fc326Eaad87f3B876b3A5a7db245F3c"),
		avsOwnerAddress,
		assetIds,
		uint64(avsUnbondingPeriod),
		uint64(minSelfDelegation),
		epochIdentifier,
		params,
	)
	log.Println(err)

}

func TestRunCreateTask(t *testing.T) {
	avsAddr, _, err := DeployAVS()
	//The successful execution of the transaction for deploying the contract
	//and packaging it into the block requires waiting for a moment
	time.Sleep(5 * time.Second)
	log.Println("avs###" + avsAddr.String())
	service, _ := NewExoClientService(avsAddr)
	log.Println(service.ethClient.ChainID(context.Background()))
	avsName, epochIdentifier := "avsTest", "hour"
	avsOwnerAddress := []string{"exo13h6xg79g82e2g2vhjwg7j4r2z2hlncelwutkjr", "exo1sc9kjykz6qehauzmhjympsktdjaw4d99dksgrk"}
	assetIds := []string{"0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"}
	minStakeAmount, avsUnbondingPeriod, minSelfDelegation := 3, 3, 5
	params := []uint64{5, 7, 8, 4}
	_, err = service.exocoreWriter.RegisterAVSToExocore(context.Background(),
		avsName,
		uint64(minStakeAmount),
		avsAddr,
		gethcommon.HexToAddress("0x92D203486fc326Eaad87f3B876b3A5a7db245F3c"),
		gethcommon.HexToAddress("0x92D203486fc326Eaad87f3B876b3A5a7db245F3c"),
		avsOwnerAddress,
		assetIds,
		uint64(avsUnbondingPeriod),
		uint64(minSelfDelegation),
		epochIdentifier,
		params,
	)
	log.Println(err)
	time.Sleep(5 * time.Second)
	service.exocoreWriter.CreateNewTask(
		context.Background(),
		"hello-world",
		uint64(5),
		uint64(5),
		uint64(90),
		uint64(5),
	)
	log.Println(err)
}

func TestRunSubTask(t *testing.T) {
	avsAddr, _, _ := DeployAVS()
	//The successful execution of the transaction for deploying the contract
	//and packaging it into the block requires waiting for a moment
	time.Sleep(5 * time.Second)
	log.Println("avs###" + avsAddr.String())
	service, _ := NewExoClientService(avsAddr)
	log.Println(service.ethClient.ChainID(context.Background()))
	newTaskCreatedChan := make(chan *avssub.ContractavsserviceTaskCreated)
	sub := service.exocoreSub.SubscribeToNewTasks(newTaskCreatedChan)
	i := 1
	for {
		log.Println(i)
		i++
		select {
		case <-context.Background().Done():
			log.Println("done")
		case err := <-sub.Err():
			log.Println("Error in websocket subscription", "err", err)
			sub.Unsubscribe()
			sub = service.exocoreSub.SubscribeToNewTasks(newTaskCreatedChan)
		case newTaskCreatedLog := <-newTaskCreatedChan:
			log.Println("newTaskCreatedChan：", newTaskCreatedLog)
		}
	}
}

func TestGetCode(t *testing.T) {

	client, err := ethclient.Dial(EthUrl)
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress("0x5708eF2ec22A306A4F5F556052507352040072A0")

	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		panic(err)
	}

	log.Printf("Contract code = %v\n", code)

}

func TestGetBalance(t *testing.T) {

	client, err := ethclient.Dial(EthUrl)
	if err != nil {
		panic(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("Cannot get chainId", "err", err)
		return
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Println("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
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
	log.Println(senderAddress.String())
	log.Printf("balance = %v\n", balance)

}

func Test_eth_subscribe(t *testing.T) {

	client, err := eth.NewClient(EthWsUrl)
	if err != nil {
		log.Println("Cannot create ws ethclient", "err", err)
	}

	contractAddress := common.HexToAddress("0x92D203486fc326Eaad87f3B876b3A5a7db245F3c")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case vLog := <-logs:
			log.Println(vLog) // pointer to event log
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
		log.Println("pri:", hex.EncodeToString(pri))
		log.Println("pub:", hex.EncodeToString(pub))

		privateKeys[i] = privateKey
		publicKeys[i] = privateKey.PublicKey()
		log.Println(len(privateKey.PublicKey().Marshal()))
		sigs[i] = privateKey.Sign(msg[:])
		dsig := privateKey.Sign(msg[:]).Marshal()
		log.Println(len(dsig))

		log.Println("dsig:", hex.EncodeToString(dsig))
		valid := privateKey.Sign(msg[:]).Verify(privateKey.PublicKey(), msg[:])
		require.True(t, valid, "Signature verification failed")

	}
	// s.publicKeys = append(s.publicKeys[:1], s.publicKeys[1+1:]...)

	aggregatedPublicKey := blst.AggregateMultiplePubkeys(publicKeys)
	aggregatedPub := aggregatedPublicKey.Marshal()
	log.Println("aggregatedPublicKey:", hex.EncodeToString(aggregatedPub))

	signature := blst.AggregateSignatures(sigs)

	valid := signature.FastAggregateVerify(publicKeys, msg)
	valid1 := signature.Verify(aggregatedPublicKey, msg.Bytes())

	require.True(t, valid, "Signature verification failed")
	require.True(t, valid1, "Signature verification failed")

	log.Println("msg:" + msg.Hex())
	aggsig := signature.Marshal()
	log.Println("aggsignature:", hex.EncodeToString(aggsig)) // ok

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
	log.Println("Aggregate signature1 is valid for all messages:", valid1)
	log.Println("Aggregate signature2 is valid for all messages:", valid2)
	log.Println("Aggregate signature3 is valid for all messages:", valid3)

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
	log.Println("Aggregate signature2 is valid for all messages:", valid2)
	log.Println("Aggregate signature3 is valid for all messages:", valid3)

	sigN := privateKeys[1].Sign(msg.Bytes())
	a := sigN.Marshal()
	log.Println(a)
	valid := sigN.Verify(publicKeys[1], msg.Bytes())
	log.Println(" sigN is valid for all messages:", valid)

	b, _ := blst.VerifySignature(a, [32]byte(msg.Bytes()), publicKeys[1])
	log.Println(" b is valid for all messages:", b)

}
