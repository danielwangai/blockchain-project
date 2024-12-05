package node

import (
	"encoding/hex"
	"fmt"
	"github.com/danielwangai/blockchain-project/proto"
	"github.com/danielwangai/blockchain-project/types"
)

type HeaderList struct {
	Headers []*proto.Header
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		Headers: []*proto.Header{},
	}
}

func (hl *HeaderList) Add(header *proto.Header) {
	hl.Headers = append(hl.Headers, header)
}

// Len returns the length of the header list
func (hl *HeaderList) Len() int {
	return len(hl.Headers)
}

// Height = length - 1 since the first block in the chain(genesis block)
// starts from 0
func (hl *HeaderList) Height() int {
	return hl.Len() - 1
}

func (hl *HeaderList) Get(index int) *proto.Header {
	if index < 0 || index >= len(hl.Headers) {
		panic("height too high")
	}

	return hl.Headers[index]
}

type Chain struct {
	headers    *HeaderList
	blockStore BlockStorer
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{blockStore: bs, headers: NewHeaderList()}
}

func (c *Chain) Height() int {
	return c.headers.Height()
}

func (c *Chain) AddBlock(b *proto.Block) error {
	// append header to header list
	c.headers.Add(b.Header)
	return c.blockStore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.blockStore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if height < c.Height() {
		return nil, fmt.Errorf("height too high. block height [%d], chain height [%d]", height, c.Height())
	}
	header := c.headers.Get(height)
	hash := types.HashHeader(header)
	return c.GetBlockByHash(hash)
}
