package types

var (
	prefixCfg = &AddrPrefixCfg{
		AccountAddr:   "exo",
		ValidatorAddr: "exovaloper",
		ConsensusAddr: "exovalcons",
		AccountPub:    "exopub",
		ValidatorPub:  "exovaloperpub",
		ConsensusPub:  "exovalconspub",
	}
)

type AddrPrefixCfg struct {
	AccountAddr   string
	ValidatorAddr string
	ConsensusAddr string
	AccountPub    string
	ValidatorPub  string
	ConsensusPub  string
}

func setAddrPrefix(prefix *AddrPrefixCfg) {
	prefixCfg = prefix
}

// GetAddrPrefixCfg returns the config instance for the corresponding Network type
func GetAddrPrefixCfg() *AddrPrefixCfg {
	return prefixCfg
}

// GetBech32AccountAddrPrefix returns the Bech32 prefix for account address
func (config *AddrPrefixCfg) GetBech32AccountAddrPrefix() string {
	return config.AccountAddr
}

// GetBech32ValidatorAddrPrefix returns the Bech32 prefix for validator address
func (config *AddrPrefixCfg) GetBech32ValidatorAddrPrefix() string {
	return config.ValidatorAddr
}

// GetBech32ConsensusAddrPrefix returns the Bech32 prefix for consensus node address
func (config *AddrPrefixCfg) GetBech32ConsensusAddrPrefix() string {
	return config.ConsensusAddr
}

// GetBech32AccountPubPrefix returns the Bech32 prefix for account public key
func (config *AddrPrefixCfg) GetBech32AccountPubPrefix() string {
	return config.AccountPub
}

// GetBech32ValidatorPubPrefix returns the Bech32 prefix for validator public key
func (config *AddrPrefixCfg) GetBech32ValidatorPubPrefix() string {
	return config.ValidatorPub
}

// GetBech32ConsensusPubPrefix returns the Bech32 prefix for consensus node public key
func (config *AddrPrefixCfg) GetBech32ConsensusPubPrefix() string {
	return config.ConsensusPub
}
