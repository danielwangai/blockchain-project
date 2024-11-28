package types

import (
	"crypto/sha256"
	"github.com/danielwangai/blockchain-project/crypto"
	"github.com/danielwangai/blockchain-project/proto"
	pb "github.com/golang/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

// VerifyTransaction verifies input(s) only
// assumes that output was verified earlier as it was an input of
// a previous transaction
func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		signature := crypto.SignatureFromBytes(input.Signature)
		pubKey := crypto.PublicKeyFromBytes(input.PublicKey)
		// TODO: remove signature from the input
		input.Signature = nil
		if !signature.Verify(pubKey, HashTransaction(tx)) {
			return false
		}
	}

	return true
}
