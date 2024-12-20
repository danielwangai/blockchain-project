package node

import (
	"github.com/danielwangai/blockchain-project/types"
	"github.com/danielwangai/blockchain-project/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	// create a block
	b := utils.RandomBlock()
	hash := types.HashBlock(b)
	// add block to chain, should not return error
	assert.Nil(t, chain.AddBlock(b))
	// fetch created block
	fetchedBlock, err := chain.GetBlockByHash(hash)
	assert.Nil(t, err)
	// compare to created block
	assert.Equal(t, b, fetchedBlock)
}

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	// height of -1 since there are 0 blocks in the chain
	assert.Equal(t, -1, chain.Height())
	block := utils.RandomBlock()
	chain.AddBlock(block)
	assert.Equal(t, 0, chain.Height())

	// test get block by height
	blockByHeight, err := chain.GetBlockByHeight(0)
	assert.Nil(t, err)
	assert.Equal(t, block, blockByHeight)
}
