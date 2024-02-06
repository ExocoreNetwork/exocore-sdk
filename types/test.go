package types

import "github.com/ExocoreNetwork/exocore-sdk/crypto/bls"

type TestOperator struct {
	OperatorId     OperatorId
	StakePerQuorum map[QuorumNum]StakeAmount
	BlsKeypair     *bls.KeyPair
}
