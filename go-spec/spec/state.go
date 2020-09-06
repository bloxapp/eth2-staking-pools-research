package spec

import (
	"encoding/hex"
	"fmt"
	"github.com/prysmaticlabs/go-ssz"
)

type BlockProducer struct {
	Id				uint64
	Balance			uint64 // balance on the pool chain (rewards earned)
	Stake			uint64 // stake
	Slashed			bool
}

type Pool struct {
	Id				[]byte // pubkey
	SortedExecutors		[]uint64 // ids of the block producers which are executors on this pool
}

type State struct {
	Pools			[]*Pool
	BlockRoots		[]byte
	HeadBlockHeader	[]*BlockHeader
	BlockProducers  map[uint64]*BlockProducer
	Seed			[]byte
}

func (state *State) Copy() *State {
	return nil
}

func (state *State) Root() ([32]byte,error) {
	return ssz.HashTreeRoot(state)
}

func (state *State) IsActivePool(pool []byte) bool {
	return true // TODO
}

func (state *State) IncreaseBlockProducerBalance(bp uint64, change uint64) (newBalance uint64, error error) {
	state.BlockProducers[bp].Balance += change
	return state.BlockProducers[bp].Balance, nil
}

func (state *State) DecreaseBlockProducerBalance(bp uint64, change uint64) (newBalance uint64, error error) {
	if state.BlockProducers[bp].Balance < change {
		return 0, fmt.Errorf("BP %d dosen't have enonugh balance (%d) to decrease (%d)", bp, state.BlockProducers[bp].Balance, change)
	}

	state.BlockProducers[bp].Balance -= change
	return state.BlockProducers[bp].Balance, nil
}

// Applies every pool performance to its relevant executors, decreasing and increasing balances.
func (state *State) ApplyPoolExecutions(summaries []*PoolExecutionSummary) error {
	for _, summary := range summaries {
		if !state.IsActivePool(summary.PoolId) {
			return fmt.Errorf("pool %s not active", hex.EncodeToString(summary.PoolId))
		}

		if err := summary.ApplyOnState(state); err != nil {
			return err
		}
	}
	return nil
}

// called when a new block was proposed
func (state *State) ProcessNewBlock(newBlockHeader *BlockHeader) (newState *State, error error) {
	newBlock,err := GetBlockBody(newBlockHeader.BlockRoot)
	if err != nil {
		return nil, err
	}

	// copy the state to apply state transition on
	stateCopy := state.Copy()

	err = stateCopy.ApplyPoolExecutions(newBlock.PoolsExecutionSummary)
	if err != nil {
		return nil, err
	}

	return stateCopy, nil
}