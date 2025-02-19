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

	expectedCandidates := [][]obxor.Candidate{
		{ // first byte of the key
			{ B: 0, Result: []byte{1,3} },
			{ B: 2, Result: []byte{3,1} },
		},
		{ // second byte of the key
			{ B: 0, Result: []byte{2,1} },
		},
	}
	
	res, err := x.Candidates()
	assert.DeepEqual(t, res, expectedCandidates)
	assert.Equal(t, err, nil)
}
