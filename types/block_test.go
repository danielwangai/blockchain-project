package types

import (
	"github.com/danielwangai/blockchain-project/crypto"
	"github.com/danielwangai/blockchain-project/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashBlock(t *testing.T) {
	block := utils.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}

func TestSignBlock(t *testing.T) {
	pk := crypto.GeneratePrivateKey()
	block := utils.RandomBlock()

	pubKey := pk.Public()
	signature := SignBlock(pk, block)
	assert.Equal(t, 64, len(signature.Bytes()))
	assert.True(t, signature.Verify(pubKey, HashBlock(block)))
}
