package state_transition

import (
	"github.com/bloxapp/eth2-staking-pools-research/go-spec/src/core"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/prysmaticlabs/go-ssz"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStateCopying(t *testing.T) {
	require.NoError(t, bls.Init(bls.BLS12_381))
	require.NoError(t, bls.SetETHmode(bls.EthModeDraft07))

	state := generateTestState(t)

	preRoot, err := ssz.HashTreeRoot(state)
	require.NoError(t, err)

	newState := core.CopyState(state)
	require.NoError(t, err)

	// test new state and old state ssz
	newStateRoot, err := ssz.HashTreeRoot(newState)
	require.NoError(t, err)
	require.EqualValues(t, preRoot, newStateRoot)

	// test manipulating prams on new state copying
	bp := core.GetBlockProducer(newState, 0)
	bp.Balance = 100000
	require.NotEqualValues(t, core.GetBlockProducer(state, 0).Balance, 100000)

	core.GetPool(newState, 0).Active = false
	require.NotEqualValues(t, core.GetPool(state, 0).Active, core.GetPool(newState, 0).Active)

	// test old state root not changed
	postRoot, err := ssz.HashTreeRoot(state)
	require.NoError(t, err)
	require.EqualValues(t, preRoot, postRoot)
}

func TestBlockApplyConsistency(t *testing.T) {
	require.NoError(t, bls.Init(bls.BLS12_381))
	require.NoError(t, bls.SetETHmode(bls.EthModeDraft07))

	state := generateTestState(t)
	_, body := GenerateValidHeadAndBody(state)

	preRoot,err := ssz.HashTreeRoot(state)
	require.NoError(t, err)

	var postRoot []byte
	for i := 0 ; i < 10 ; i++ {
		st := NewStateTransition()
		newState, err := st.ApplyBlock(state, body)
		require.NoError(t, err)

		if i != 0 {
			require.EqualValues(t, postRoot, core.GetStateRoot(newState, 1))
		}

		postRoot = core.GetStateRoot(newState, 1)
	}

	post,err := ssz.HashTreeRoot(state)
	require.NoError(t, err)
	require.EqualValues(t, preRoot, post)
}

func TestBlockApplyPostRoot(t *testing.T) {
	t.Skipf("need to test static state transition")
	// TODO - need to test static state transition
	//require.NoError(t, bls.Init(bls.BLS12_381))
	//require.NoError(t, bls.SetETHmode(bls.EthModeDraft07))
	//
	//state := generateTestState(t)
	//_, body := GenerateValidHeadAndBody(state)
	//st := NewStateTransition()
	//newState, err := st.ApplyBlock(state, body)
	//require.NoError(t, err)
	//
	//fmt.Printf("%s\n", hex.EncodeToString(core.GetStateRoot(newState, newState.CurrentEpoch)))
	//require.EqualValues(t, toByte(""), core.GetStateRoot(newState, newState.CurrentEpoch))
}