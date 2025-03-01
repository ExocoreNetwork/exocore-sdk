package bls

// TODO: put in proper module later

import (
	"context"
	"math/big"
)

// Operators

type StakeAmount *big.Int

// OperatorInfo contains information about an operator which is stored on the blockchain state,
// corresponding to a particular quorum
type OperatorInfo struct {
	// Stake is the amount of stake held by the operator in the quorum
	Stake StakeAmount
	// Index is the index of the operator within the quorum
	Index OperatorIndex
}

// OperatorState contains information about the current state of operators which is stored in the blockchain state
type OperatorState struct {
	// Operators is a map from quorum ID to a map from the operators in that quorum to their StoredOperatorInfo.
	// Membership
	// in the map implies membership in the quorum.
	Operators map[QuorumID]map[OperatorId]*OperatorInfo
	// Totals is a map from quorum ID to the total stake (Stake) and total count (Index) of all operators in that quorum
	Totals map[QuorumID]*OperatorInfo
	// BlockNumber is the block number at which this state was retrieved
	BlockNumber uint
}

// IndexedOperatorInfo contains information about an operator which is contained in events from the avs smart
// contracts. Note that
// this information does not depend on the quorum.
type IndexedOperatorInfo struct {
	// PubKeyG1 and PubKeyG2 are the public keys of the operator, which are retrieved from the BlsApkRegistry
	// smart contract
	PubkeyG1 *G1Point
	PubkeyG2 *G2Point
	// Socket is the socket address of the operator, in the form "host:port"
	Socket string
}

// IndexedOperatorState contains information about the current state of operators which is contained in events from the
// in addition to the information contained in OperatorState
type IndexedOperatorState struct {
	*OperatorState
	// IndexedOperators is a map from operator ID to the IndexedOperatorInfo for that operator.
	IndexedOperators map[OperatorId]*IndexedOperatorInfo
	// AggKeys is a map from quorum ID to the aggregate public key of the operators in that quorum
	AggKeys map[QuorumID]*G1Point
}

// ChainState is an interface for getting information about the current chain state.
type ChainState interface {
	GetCurrentBlockNumber() (uint, error)
	GetOperatorState(blockNumber uint, quorums []QuorumID) (*OperatorState, error)
	GetOperatorStateByOperator(blockNumber uint, operator OperatorId) (*OperatorState, error)
	// GetOperatorQuorums(blockNumber uint, operator OperatorId) ([]uint, error)
}

// ChainState is an interface for getting information about the current chain state.
type IndexedChainState interface {
	ChainState
	GetIndexedOperatorState(blockNumber uint, quorums []QuorumID) (*IndexedOperatorState, error)
	Start(context context.Context) error
}
