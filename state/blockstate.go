package state

import (
	"sync"

	"github.com/aergoio/aergo/consensus"
	"github.com/aergoio/aergo/types"
	"github.com/willf/bloom"
)

// BlockInfo contains BlockHash and StateRoot
type BlockInfo struct {
	BlockHash types.BlockID
	StateRoot types.HashID
}

// BlockState contains BlockInfo and statedb for block
type BlockState struct {
	StateDB
	BpReward      []byte //final bp reward, increment when tx executes
	receipts      types.Receipts
	CodeMap       codeCache
	CCProposal    *consensus.ConfChangePropose
	prevBlockHash []byte
}

type codeCache struct {
	Lock  sync.Mutex
	codes map[types.AccountID][]byte
}

// NewBlockInfo create new blockInfo contains blockNo, blockHash and blockHash of previous block
func NewBlockInfo(blockHash types.BlockID, stateRoot types.HashID) *BlockInfo {
	return &BlockInfo{
		BlockHash: blockHash,
		StateRoot: stateRoot,
	}
}

// GetStateRoot return bytes of bi.StateRoot
func (bi *BlockInfo) GetStateRoot() []byte {
	if bi == nil {
		return nil
	}
	return bi.StateRoot.Bytes()
}

type BlockStateOptFn func(s *BlockState)

func SetPrevBlockHash(h []byte) BlockStateOptFn {
	return func(s *BlockState) {
		s.SetPrevBlockHash(h)
	}
}

// NewBlockState create new blockState contains blockInfo, account states and undo states
func NewBlockState(states *StateDB, options ...BlockStateOptFn) *BlockState {
	b := &BlockState{
		StateDB: *states,
		CodeMap: codeCache{
			codes: make(map[types.AccountID][]byte),
		},
	}
	for _, opt := range options {
		opt(b)
	}
	return b
}

func (bs *BlockState) AddReceipt(r *types.Receipt) error {
	if len(r.Events) > 0 {
		rBloom := bloom.New(types.BloomBitBits, types.BloomHashKNum)
		for _, e := range r.Events {
			rBloom.Add(e.ContractAddress)
			rBloom.Add([]byte(e.EventName))
		}
		binary, _ := rBloom.GobEncode()
		r.Bloom = binary[24:]
		err := bs.receipts.MergeBloom(rBloom)
		if err != nil {
			return err
		}
	}
	bs.receipts.Set(append(bs.receipts.Get(), r))
	return nil
}

func (bs *BlockState) Receipts() *types.Receipts {
	if bs == nil {
		return nil
	}
	return &bs.receipts
}

func (bs *BlockState) SetPrevBlockHash(prevHash []byte) *BlockState {
	if bs != nil {
		bs.prevBlockHash = prevHash
	}
	return bs
}

func (bs *BlockState) PrevBlockHash() []byte {
	return bs.prevBlockHash
}

func (c *codeCache) Add(key types.AccountID, code []byte) {
	c.Lock.Lock()
	c.codes[key] = code
	c.Lock.Unlock()
}

func (c *codeCache) Get(key types.AccountID) []byte {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	return c.codes[key]
}

func (c *codeCache) Remove(key types.AccountID) {
	c.Lock.Lock()
	delete(c.codes, key)
	c.Lock.Unlock()
}
