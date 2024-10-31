package generate_test

import (
	"encoding/hex"
	"fmt"
	sdkBls "github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	sdkEcdsa "github.com/ExocoreNetwork/exocore-sdk/crypto/ecdsa"

	"github.com/prysmaticlabs/prysm/v4/crypto/bls/blst"
	"os"
	"path/filepath"
	"testing"

	key "github.com/ExocoreNetwork/exocore-sdk/cmd/exokey/generate"
)

const (
	KeystorePath = "/keys/key.json"
)

func TestEcdsaKey(t *testing.T) {

	ecdsaKeyPassword, _ := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")

	expectedPrivateKey := "d196dca836f8ac2fff45b3c9f0113825ccbb33fa1b39737b948503b263ed75ae"

	err := key.ImportEcdsaKeyToFile(expectedPrivateKey, ecdsaKeyPassword)
	if err != nil {
		t.Fatalf("Error importing key: %v", err)
	}

	currentDir, _ := os.Getwd()
	KeystorePath := filepath.Join(currentDir, KeystorePath)

	importedKey, err := sdkEcdsa.ReadKey(KeystorePath, ecdsaKeyPassword)
	if err != nil {
		t.Fatalf("Error reading imported key: %v", err)
	}

	actualPrivateKey := hex.EncodeToString(importedKey.D.Bytes())
	if actualPrivateKey != expectedPrivateKey {
		t.Fatalf("Expected private key: %s, but got: %s", expectedPrivateKey, actualPrivateKey)
	}

	fmt.Println("Test passed: Imported and read ecdsa key successfully")
}

func TestBlsKey(t *testing.T) {

	blsKeyPassword, _ := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	privateKey, err := blst.RandKey()
	expectedPrivateKey := hex.EncodeToString(privateKey.Marshal())
	currentDir, _ := os.Getwd()

	KeystorePath := filepath.Join(currentDir, KeystorePath)

	err = key.SaveToFile(privateKey, KeystorePath, blsKeyPassword)
	if err != nil {
		t.Fatalf("Error importing key: %v", err)
	}

	importedKey, err := sdkBls.ReadPrivateKeyFromFile(KeystorePath, blsKeyPassword)
	if err != nil {
		t.Fatalf("Error reading imported key: %v", err)
	}

	actualPrivateKey := hex.EncodeToString(importedKey.Marshal())
	if actualPrivateKey != expectedPrivateKey {
		t.Fatalf("Expected private key: %s, but got: %s", expectedPrivateKey, actualPrivateKey)
	}

	fmt.Println("Test passed: Imported and read bls key successfully")
}
