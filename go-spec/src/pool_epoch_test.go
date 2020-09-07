package src

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func GenerateAttestationSuccessfulSummary() *PoolExecutionSummary {
	return &PoolExecutionSummary{
		PoolId:        0,
		StartingEpoch: 0,
		EndEpoch:      1,
		Performance:   map[*BeaconDuty][16]byte{
			&BeaconDuty{
				Type:     0,
				Slot:     0,
				Included: true,
			}: [16]byte{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, // the first executor index is set to 1
		},
	}
}

func TestAttestationSuccessful(t *testing.T) {
	state := GenerateRandomState()
	summary := GenerateAttestationSuccessfulSummary()

	require.NoError(t, summary.ApplyOnState(state))

	for _, whoExecuted := range summary.Performance {
		pool, err := GetPool(state, summary.PoolId)
		require.NoError(t, err)

		for i:=0 ; i < int(TestConfig().PoolExecutorsNumber) ; i++ {
			bp,err := GetBlockProducer(state, pool.SortedExecutors[i])
			require.NoError(t, err)

			if IsBitSet(whoExecuted[:], uint64(i)) {
				require.EqualValues(t, 1100, bp.Balance)
				require.EqualValues(t, 0, i)
			} else {
				require.EqualValues(t, 900, bp.Balance)
			}
		}
	}
}