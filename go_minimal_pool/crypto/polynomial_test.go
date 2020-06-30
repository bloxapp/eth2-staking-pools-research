package crypto

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/stretchr/testify/require"
	"testing"
)

func frFromInt(i int64) bls.Fr {
	p := bls.Fr{}
	p.SetInt64(i)
	return p
}

func TestEvaluation(t *testing.T) {
	initBLS()

	sk := bls.Fr{}
	sk.SetByCSPRNG()
	degree := uint8(3)

	p, err := NewPolynomial(sk, degree)
	require.NoError(t,err)
	err = p.GenerateRandom()
	require.NoError(t,err)

	//fmt.Printf("%s\n", p.toString())

	res1, err := p.Evaluate(1)
	require.NoError(t,err)
	res2, err := p.Evaluate(2)
	require.NoError(t,err)
	res3, err := p.Evaluate(3)
	require.NoError(t,err)
	res4, err := p.Evaluate(4)
	require.NoError(t,err)

	// interpolate back
	points := [][]bls.Fr {
		{frFromInt(1), *res1},
		{frFromInt(2), *res2},
		{frFromInt(3), *res3},
		{frFromInt(4), *res4},
	}
	pInter := NewLagrangeInterpolation(points)
	res, err := pInter.interpolate()
	require.NoError(t,err)

	require.Equal(t, sk.GetString(10), res.GetString(10))
}

func TestInterpolation(t *testing.T) {
	initBLS()

	points := [][]bls.Fr {
		{frFromInt(1), frFromInt(7)},
		{frFromInt(2), frFromInt(10)},
		{frFromInt(3), frFromInt(15)},
	}

	p := NewLagrangeInterpolation(points)
	res, err := p.interpolate()
	require.NoError(t,err)

	require.Equal(t, "6", res.GetString(10))
}