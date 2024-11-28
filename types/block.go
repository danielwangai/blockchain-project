package types

import (
	"crypto/sha256"
	"github.com/danielwangai/blockchain-project/crypto"
	"github.com/danielwangai/blockchain-project/proto"
	pb "github.com/golang/protobuf/proto"
)

// HashBlock creates a SHA256 of the Header
func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	// sha256.Sum256(b) returns an array
	return hash[:] // returns a slice
}

func SignBlock(pk *crypto.PrivateKey, block *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(block))
}
