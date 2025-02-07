package types

import (
	"crypto/x509"
	"fmt"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v16/crypto/ethsecp256k1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultGas           = 200000
	defaultFees          = "4aexo"
	defaultTimeout       = 5
	defaultLevel         = "info"
	defaultMaxTxsBytes   = 1073741824
	defaultAlgo          = "secp256k1"
	defaultMode          = Sync
	defaultGasAdjustment = 1.0
	defaultTxSizeLimit   = 1048576
	BIP44Prefix          = "44'/118'/"
	PartialPath          = "0'/0/0"
	FullPath             = "m/" + BIP44Prefix + PartialPath
)
const (
	// maxMemoCharacters = 100
	// txSigLimit        = 7
	maxGasWanted = uint64((1 << 63) - 1)

	Sync string = "sync"
)

type ClientConfig struct {
	// RPCAddr node rpc address
	RPCAddr string

	// aexohub grpc address
	GRPCAddr string

	// grpc dial options
	GRPCOptions []grpc.DialOption

	// aexohub chain-id
	ChainID string

	// max gas limit
	Gas uint64

	// Private key generation algorithm(sm2,secp256k1)
	Algo string

	// Timeout for accessing the blockchain (such as query transactions, broadcast transactions, etc.)
	Timeout uint

	// log level(trace|debug|info|warn|error|fatal|panic)
	Level string

	// maximum bytes of a transaction
	MaxTxBytes uint64

	// adjustment factor to be multiplied against the estimate returned by the tx simulation;
	GasAdjustment float64

	// whether to enable caching
	Cached bool

	TxSizeLimit uint64

	// BIP44 path
	BIP44Path string

	// Header for rpc or http
	Header map[string][]string

	// WSAddr for ws or wss protocol
	WSAddr string

	KeyStr   string
	Priv     *ethsecp256k1.PrivKey
	GasPrice uint64
	// bech32 Address Prefix
	bech32Prefix *AddrPrefixCfg

	FeeGranter csdk.AccAddress
	FeePayer   csdk.AccAddress
}

func NewClientConfig(rpcAddr, grpcAddr, chainID, KeyStr string, gas, gasPrice uint64, options ...Option) (ClientConfig, error) {
	cfg := ClientConfig{
		RPCAddr:  rpcAddr,
		ChainID:  chainID,
		GRPCAddr: grpcAddr,
		KeyStr:   KeyStr,
		GasPrice: gasPrice,
		Gas:      gas,
	}
	for _, optionFn := range options {
		if err := optionFn(&cfg); err != nil {
			return ClientConfig{}, err
		}
	}

	if err := cfg.checkAndSetDefault(); err != nil {
		return ClientConfig{}, err
	}
	return cfg, nil
}

func (cfg *ClientConfig) checkAndSetDefault() error {
	if len(cfg.RPCAddr) == 0 {
		return fmt.Errorf("nodeURI is required")
	}

	if len(cfg.ChainID) == 0 {
		return fmt.Errorf("chainID is required")
	}

	if err := GasOption(cfg.Gas)(cfg); err != nil {
		return err
	}

	if err := AlgoOption(cfg.Algo)(cfg); err != nil {
		return err
	}

	if err := TimeoutOption(cfg.Timeout)(cfg); err != nil {
		return err
	}

	if err := LevelOption(cfg.Level)(cfg); err != nil {
		return err
	}

	if err := MaxTxBytesOption(cfg.MaxTxBytes)(cfg); err != nil {
		return err
	}

	if err := TxSizeLimitOption(cfg.TxSizeLimit)(cfg); err != nil {
		return err
	}

	if err := Bech32AddressPrefixOption(cfg.bech32Prefix)(cfg); err != nil {
		return err
	}

	if err := BIP44PathOption(cfg.BIP44Path)(cfg); err != nil {
		return err
	}
	return GasAdjustmentOption(cfg.GasAdjustment)(cfg)
}

type Option func(cfg *ClientConfig) error

func GasOption(gas uint64) Option {
	return func(cfg *ClientConfig) error {
		if gas <= 0 {
			gas = defaultGas
		}
		cfg.Gas = gas
		return nil
	}
}

func AlgoOption(algo string) Option {
	return func(cfg *ClientConfig) error {
		if algo == "" {
			algo = defaultAlgo
		}
		cfg.Algo = algo
		return nil
	}
}

func TimeoutOption(timeout uint) Option {
	return func(cfg *ClientConfig) error {
		if timeout <= 0 {
			timeout = defaultTimeout
		}
		cfg.Timeout = timeout
		return nil
	}
}

func LevelOption(level string) Option {
	return func(cfg *ClientConfig) error {
		if level == "" {
			level = defaultLevel
		}
		cfg.Level = level
		return nil
	}
}

func MaxTxBytesOption(maxTxBytes uint64) Option {
	return func(cfg *ClientConfig) error {
		if maxTxBytes <= 0 {
			maxTxBytes = defaultMaxTxsBytes
		}
		cfg.MaxTxBytes = maxTxBytes
		return nil
	}
}

func GasAdjustmentOption(gasAdjustment float64) Option {
	return func(cfg *ClientConfig) error {
		if gasAdjustment <= 0 {
			gasAdjustment = defaultGasAdjustment
		}
		cfg.GasAdjustment = gasAdjustment
		return nil
	}
}

func CachedOption(enabled bool) Option {
	return func(cfg *ClientConfig) error {
		cfg.Cached = enabled
		return nil
	}
}

func TxSizeLimitOption(txSizeLimit uint64) Option {
	return func(cfg *ClientConfig) error {
		if txSizeLimit <= 0 {
			txSizeLimit = defaultTxSizeLimit
		}
		cfg.TxSizeLimit = txSizeLimit
		return nil
	}
}

func Bech32AddressPrefixOption(prefix *AddrPrefixCfg) Option {
	return func(cfg *ClientConfig) error {
		if prefix != nil {
			setAddrPrefix(prefix)
		}
		return nil
	}
}

func BIP44PathOption(bIP44Path string) Option {
	return func(cfg *ClientConfig) error {
		if bIP44Path == "" {
			bIP44Path = FullPath
		}
		cfg.BIP44Path = bIP44Path
		return nil
	}
}

func HeaderOption(header map[string][]string) Option {
	return func(cfg *ClientConfig) error {
		cfg.Header = header
		return nil
	}
}

func WSAddrOption(wsAddr string) Option {
	return func(cfg *ClientConfig) error {
		cfg.WSAddr = wsAddr
		return nil
	}
}

func GRPCOptions(gRPCOptions []grpc.DialOption, TLS bool, rpcAddr string) Option {
	return func(cfg *ClientConfig) error {
		if !TLS {
			cfg.GRPCOptions = gRPCOptions
			return nil
		}

		certificateList, err := GetTLSCertPool(rpcAddr)
		if err != nil {
			panic(err)
		}

		roots := x509.NewCertPool()
		for i := range certificateList {
			roots.AddCert(certificateList[i])
		}
		cert := credentials.NewClientTLSFromCert(roots, "")
		cfg.GRPCOptions = append(gRPCOptions, grpc.WithTransportCredentials(cert))

		return nil
	}
}
