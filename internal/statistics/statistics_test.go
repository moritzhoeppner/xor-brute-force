package statistics

import (
	"testing"
	"gotest.tools/v3/assert"
)

func TestRelByteDist(t *testing.T) {
	relDist := RelByteDist([]byte{ 1, 2, 1, 4 })
	assert.DeepEqual(t, relDist, map[byte]float64{ 1: 0.5, 2: 0.25, 4: 0.25 })
}

func TestChiSquarePerfectFit(t *testing.T) {
	dist := map[byte]float64{
		1: 4.0,
		2: 7.0,
		3: 1.0,
	}
	assert.Equal(t, ChiSquare(dist, dist), 0.0)
}

func TestChiSquareImPerfectFit(t *testing.T) {
	observed := map[byte]float64{
		1: 4.0,
		2: 7.0,
		3: 1.0,
	}
	theoretical := map[byte]float64{
		1: 2.0,
		2: 8.0,
		3: 1.0,
	}

	assert.Equal(t, ChiSquare(observed, theoretical), 2.125)
}

func TestChiSquareMissingValue(t *testing.T) {
	observed := map[byte]float64{
		1: 4.0,
	}
	theoretical := map[byte]float64{
		1: 4.0,
		2: 7.0,
	}
	// observed[2] is missing and should be treated as 0.
	assert.Equal(t, ChiSquare(observed, theoretical), 7.0)
}
