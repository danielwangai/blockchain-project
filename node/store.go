package node

import (
	"encoding/hex"
	"fmt"
	"github.com/danielwangai/blockchain-project/proto"
	"github.com/danielwangai/blockchain-project/types"
	"sync"
)

type BlockStorer interface {
	Put(block *proto.Block) error
	Get(string2 string) (*proto.Block, error)
}

type MemoryBlockStore struct {
	lock   sync.RWMutex
	blocks map[string]*proto.Block
}

func NewMemoryBlockStore() *MemoryBlockStore {
	return &MemoryBlockStore{blocks: make(map[string]*proto.Block)}
}

// Put adds a new block to the blockchain
func (s *MemoryBlockStore) Put(b *proto.Block) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	hash := hex.EncodeToString(types.HashBlock(b))
	s.blocks[hash] = b
	return nil
}

// Get fetches a block from the blockchain
func (s *MemoryBlockStore) Get(hash string) (*proto.Block, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	block, ok := s.blocks[hash]
	if !ok {
		return nil, fmt.Errorf("block with hash [%s] does not exist", hash)
	}

	return block, nil
}
