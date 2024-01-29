package avsregistry

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type AvsRegistryWriter interface {
	RegisterOperatorWithAVSRegistryCoordinator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey blsregistrycoordinator.BN254G1Point,
		socket string,
	) (*types.Receipt, error)
	UpdateStakes(
		ctx context.Context,
		operators []gethcommon.Address,
	) (*types.Receipt, error)

	DeregisterOperator(
		ctx context.Context,
		quorumNumbers []byte,
		pubkey blsregistrycoordinator.BN254G1Point,
	) (*types.Receipt, error)
}
