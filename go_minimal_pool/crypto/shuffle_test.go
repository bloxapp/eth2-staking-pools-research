package crypto

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)


func getSeed(str string) [32]byte {
	var ret [32]byte
	_seed, _ := hex.DecodeString(str)
	copy(ret[:],_seed)

	return ret
}

func TestSeedMix(t *testing.T) {
	tests := []struct {
		testName string
		seed [32]byte
		epoch uint32
		expected string
	}{
		{
			testName: "epoch 1",
			seed: getSeed("b581262ce281d1e9deaf2f0158d7cd05217f1196d95956c5f55d837ccc3c8a9"),
			epoch:1,
			expected: "ddffebc6e10e909165778df614d648b50b81b8647b79121da2c1142a1bb5b6a6",
		},
		{
			testName: "epoch 2",
			seed: getSeed("b581262ce281d1e9deaf2f0158d7cd05217f1196d95956c5f55d837ccc3c8a9"),
			epoch:2,
			expected: "f536fd5464af265f824e9a62144e69ecc5ef0749e5be6743dd69e28b2362e6c4",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func (t *testing.T) {
			res,err := MixSeed(test.seed, test.epoch)
			require.NoError(t,err)

			require.Equal(t, test.expected, hex.EncodeToString(res[:]))
		})
	}

}

func TestDeterministicShuffling(t *testing.T) {
	tests := []struct {
		testName string
		seed [32]byte
		rounds uint8
		indexes []uint32
		expected []uint32
	} {
		{
			testName:"shuffle, 5 rounds",
			seed: getSeed("b581262ce281d1e9deaf2f0158d7cd05217f1196d95956c5f55d837ccc3c8a9"),
			rounds: 5,
			indexes: []uint32{1,2,3,4},
			expected: []uint32{4,2,3,1},
		},
		{
			testName:"shuffle, 10 rounds",
			seed: getSeed("b581262ce281d1e9deaf2f0158d7cd05217f1196d95956c5f55d837ccc3c8a9"),
			rounds: 10,
			indexes: []uint32{1,2,3,4},
			expected: []uint32{1,2,3,4},
		},
		{
			testName:"shuffle, 15 rounds",
			seed: getSeed("b581262ce281d1e9deaf2f0158d7cd05217f1196d95956c5f55d837ccc3c8a9"),
			rounds: 15,
			indexes: []uint32{1,2,3,4},
			expected: []uint32{4,2,1,3},
		},
		{
			testName:"shuffle seed #2, 5 rounds",
			seed: getSeed("f536fd5464af265f824e9a62144e69ecc5ef0749e5be6743dd69e28b2362e6c4"),
			rounds: 5,
			indexes: []uint32{1,2,3,4},
			expected: []uint32{2,3,1,4},
		},
		{
			testName:"shuffle seed #2, 6 rounds",
			seed: getSeed("f536fd5464af265f824e9a62144e69ecc5ef0749e5be6743dd69e28b2362e6c4"),
			rounds: 6,
			indexes: []uint32{1,2,3,4},
			expected: []uint32{3,2,1,4},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := ShuffleList(test.indexes, test.seed, test.rounds)
			require.NoError(t,err)
			require.Equal(t, test.expected, res)
		})
	}
}
