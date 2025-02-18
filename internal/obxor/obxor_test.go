package obxor

import (
	"testing"
	"gotest.tools/v3/assert"
)

func TestTryByteWithAllowedResult(t *testing.T) {
	x := Obxor{
		Ciphertext: []byte{0, 3, 4, 6},
		ProbabilityDist: map[byte]float64{0: 0.4, 3: 0.25, 4: 0.1, 5: 0.1, 6: 0.05, 7: 0.1},
		KeyBytes: []byte{0, 3, 4, 6},
		ResultBytes: []byte{0, 3, 4, 5, 6, 7},
	}

	// 3 XOR ciphertext = 3075 (all in {ResultBytes})
	result, err := x.try(3)

	assert.Equal(t, err, nil)
	assert.DeepEqual(t, result, []byte{3, 0, 7, 5})
}

func TestTryByteWithNotAllowedResult(t *testing.T) {
	x := Obxor{
		Ciphertext: []byte{0, 3, 4, 6},
		ProbabilityDist: map[byte]float64{0: 0.4, 3: 0.25, 4: 0.1, 5: 0.1, 6: 0.05, 7: 0.1},
		KeyBytes: []byte{0, 3, 4, 6},
		ResultBytes: []byte{0, 3, 4, 5, 6, 7},
	}

	// 4 XOR ciphertext = 4702 (2 not in {ResultBytes})
	result, err := x.try(4)

	assert.Error(t, err, "Result contains not allowed bytes.")
	assert.DeepEqual(t, result, []byte{4, 7, 0, 2})
}

func TestCandidates(t *testing.T) {
	x := Obxor{
		Ciphertext: []byte{0, 3, 4, 6},
		ProbabilityDist: map[byte]float64{0: 0.4, 3: 0.25, 4: 0.1, 5: 0.1, 6: 0.05, 7: 0.1},
		KeyBytes: []byte{0, 3, 4, 6},
		ResultBytes: []byte{0, 3, 4, 5, 6, 7},
	}

	candidates := x.Candidates()

	// 0 XOR ciphertext = 0346 (all in {ResultBytes})
	// 3 XOR ciphertext = 3075 (all in {ResultBytes})
	// 4 XOR ciphertext = 4702 (2 not in {ResultBytes})
	// 6 XOR ciphertext = 6520 (2 not in {ResultBytes})

	assert.DeepEqual(t, candidates,
		[]Candidate{ Candidate{B: 0, Diff: 0}, Candidate{B: 3, Diff: 0} })
}
