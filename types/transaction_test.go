package types

import (
	"github.com/danielwangai/blockchain-project/crypto"
	"github.com/danielwangai/blockchain-project/proto"
	"github.com/danielwangai/blockchain-project/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTransaction(t *testing.T) {
	senderPrivKey := crypto.GeneratePrivateKey()
	senderPubKey := senderPrivKey.Public()
	senderAddr := senderPubKey.Address()

	receiverPrivKey := crypto.GeneratePrivateKey()
	receiverPubKey := receiverPrivKey.Public()
	receiverAddr := receiverPubKey.Address()

	input := &proto.TxInput{
		PrevTxHash:   utils.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    senderPubKey.Bytes(),
	}

	// send 5 units to receiverAddr
	output1 := &proto.TxOutput{
		Amount:  5,
		Address: receiverAddr.Bytes(),
	}

	// send 95 units to senderAddr
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: senderAddr.Bytes(),
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	// sign the transaction
	signature := SignTransaction(senderPrivKey, tx)
	// TODO: remove signature from the input because
	// the signature should be generated from the inputs
	// and hence can't be part of the input in the first place
	input.Signature = signature.Bytes()

	assert.Equal(t, 64, len(signature.Bytes()))
	assert.True(t, VerifyTransaction(tx))
}
