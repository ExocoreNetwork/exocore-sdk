package types

import (
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	"log/slog"
)

type TaskIndex = uint32
type TaskResponseDigest = Bytes32
type TaskResponse = interface{}

type TaskResponseHashFunction func(taskResponse TaskResponse) (TaskResponseDigest, error)

type SignedTaskResponseDigest struct {
	TaskResponse                TaskResponse
	BlsSignature                blst.Signature
	SignatureVerificationErrorC chan error `json:"-"` // removed from json because channels are not marshallable
}

func (strd SignedTaskResponseDigest) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("taskResponse", strd.TaskResponse),
		slog.Any("blsSignature", strd.BlsSignature),
	)
}
