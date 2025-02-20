package mbxor

import (
	"testing"
	"gotest.tools/v3/assert"
	"github.com/moritzhoeppner/xor-brute-force/internal/obxor"
)

func TestCandidates(t *testing.T) {
	x := Mbxor{
		Ciphertexts: [][]byte{
			{1,2},
			{3,1},
		},
		KeyBytes: []byte{0, 1, 2},
		ResultBytes: []byte{1, 2, 3},
	}

	expectedCandidates := Candidates{
		PerByte: [][]obxor.Candidate{
			{ // first byte of the key
				{ B: 0, Result: []byte{1,3} },
				{ B: 2, Result: []byte{3,1} },
			},
			{ // second byte of the key
				{ B: 0, Result: []byte{2,1} },
			},
		},
	}
	
	res, err := x.Candidates()
	assert.DeepEqual(t, res, expectedCandidates)
	assert.Equal(t, err, nil)
}

func TestMostLikely(t *testing.T) {
	x := Mbxor{
		Ciphertexts: [][]byte{
			{1,2},
			{3,1},
			{3,2},
		},
		KeyBytes: []byte{0, 1, 2},
		ResultBytes: []byte{1, 2, 3},
	}

	candidates, _ := x.Candidates()

	// Candidates for the first byte:
	//   - 0 (Result: 1,3,3)
	//   - 2 (Result: 3,1,1)
	// Candidates for the second byte:
	//	 - 0 (Result: 2,1,2)

	// If we expect more 1s then 3s, 2 the the better choice for the first byte.
	key := candidates.MostLikely(map[byte]float64 { 1: 0.5, 2: 0.25, 3: 0.25 })
	assert.DeepEqual(t, key, []byte{ 2, 0 })

	// If we expect more 3s then 1s, 0 the the better choice for the first byte.
	key = candidates.MostLikely(map[byte]float64 { 1: 0.25, 2: 0.25, 3: 0.5 })
	assert.DeepEqual(t, key, []byte{ 0, 0 })
}
