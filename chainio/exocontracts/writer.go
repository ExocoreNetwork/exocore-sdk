package exocontracts

import (
	"context"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/utils"
	blspubkeycompendium "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSPublicKeyCompendium"
	contractDelegationManager "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/DelegationManager"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signer"
	"github.com/ExocoreNetwork/exocore-sdk/types"
)

type EXOWriter interface {
	RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	UpdateOperatorDetails(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error)

	RegisterBLSPublicKey(
		ctx context.Context,
		blsKeyPair *bls.KeyPair,
		operator types.Operator,
	) (*gethtypes.Receipt, error)

	OptOperatorIntoSlashing(ctx context.Context, avsServiceManagerAddr gethcommon.Address) (*gethtypes.Receipt, error)
}

type ELChainWriter struct {
	elContractsClient clients.EXOContractsClient
	signer            signer.Signer
	logger            logging.Logger
	ethClient         eth.EthClient
}

var _ EXOWriter = (*ELChainWriter)(nil)

func NewEXOChainWriter(
	elContractsClient clients.EXOContractsClient,
	ethClient eth.EthClient,
	signer signer.Signer,
	logger logging.Logger,
) *ELChainWriter {
	return &ELChainWriter{
		elContractsClient: elContractsClient,
		signer:            signer,
		logger:            logger,
		ethClient:         ethClient,
	}
}

func (w *ELChainWriter) RegisterAsOperator(ctx context.Context, operator types.Operator) (*gethtypes.Receipt, error) {
	w.logger.Infof("registering operator %s to exocore", operator.Address)
	opDetails := contractDelegationManager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}
	txOpts := w.signer.GetTxOpts()
	tx, err := w.elContractsClient.RegisterAsOperator(txOpts, opDetails, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("registered operator %s with exocore", operator.Address)
	return receipt, nil
}

func (w *ELChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (*gethtypes.Receipt, error) {
	txOpts := w.signer.GetTxOpts()

	w.logger.Infof("updating operator details of operator %s to exocore", operator.Address)
	opDetails := contractDelegationManager.IDelegationManagerOperatorDetails{
		EarningsReceiver:         gethcommon.HexToAddress(operator.EarningsReceiverAddress),
		DelegationApprover:       gethcommon.HexToAddress(operator.DelegationApproverAddress),
		StakerOptOutWindowBlocks: operator.StakerOptOutWindowBlocks,
	}

	tx, err := w.elContractsClient.ModifyOperatorDetails(txOpts, opDetails)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("updated operator metadata URI for operator %s to exocore", operator.Address)
	tx, err = w.elContractsClient.UpdateOperatorMetadataURI(txOpts, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("updated operator details of operator %s to exocore", operator.Address)
	return receipt, nil
}

// operator opting into slashing is the w.signer wallet
// this is meant to be called by the operator CLI
func (w *ELChainWriter) OptOperatorIntoSlashing(
	ctx context.Context,
	avsServiceManagerAddr gethcommon.Address,
) (*gethtypes.Receipt, error) {
	txOpts := w.signer.GetTxOpts()
	tx, err := w.elContractsClient.OptIntoSlashing(txOpts, avsServiceManagerAddr)
	if err != nil {
		return nil, err
	}
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof(
		"Operator %s opted into slashing by service manager contract %s \n",
		txOpts.From,
		avsServiceManagerAddr,
	)
	return receipt, nil
}

func (w *ELChainWriter) RegisterBLSPublicKey(
	ctx context.Context,
	blsKeyPair *bls.KeyPair,
	operator types.Operator,
) (*gethtypes.Receipt, error) {
	w.logger.Infof("Registering BLS Public key to exocore for operator %s", operator.Address)
	txOpts := w.signer.GetTxOpts()
	chainID, err := w.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signedMsgHash := blsKeyPair.MakePubkeyRegistrationData(
		gethcommon.HexToAddress(operator.Address),
		w.elContractsClient.GetBLSPublicKeyCompendiumContractAddress(),
		chainID,
	)
	signedMsgHashBN254 := blspubkeycompendium.BN254G1Point(utils.ConvertToBN254G1Point(signedMsgHash))
	G1pubkeyBN254 := blspubkeycompendium.BN254G1Point(utils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1()))
	G2pubkeyBN254 := blspubkeycompendium.BN254G2Point(utils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2()))
	tx, err := w.elContractsClient.RegisterBLSPublicKey(
		txOpts,
		signedMsgHashBN254,
		G1pubkeyBN254,
		G2pubkeyBN254,
	)
	if err != nil {
		return nil, err
	}
	receipt := w.ethClient.WaitForTransactionReceipt(ctx, tx.Hash())

	w.logger.Infof("Operator %s has registered BLS public key to exocore \n", operator.Address)
	return receipt, nil
}
