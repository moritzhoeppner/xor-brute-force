package mbxor

import (
	"sort"
	"github.com/moritzhoeppner/xor-brute-force/internal/utils"
	"github.com/moritzhoeppner/xor-brute-force/internal/obxor"
)

type Mbxor struct {
	Ciphertexts [][]byte
	KeyBytes    []byte
	ResultBytes []byte
}

type Candidates struct {
	PerByte [][]obxor.Candidate
}

func (x *Mbxor) Candidates() (Candidates, error) {
	candidates := Candidates{}

	// Transpose Ciphertexts. The elements of the result have the same one-byte key.
	obCiphertexts, err := utils.Transpose(x.Ciphertexts)
	if err != nil {
		return candidates, err
	}

	for _, ciphertext := range obCiphertexts {
		obx := obxor.Obxor{
			Ciphertext: ciphertext,
			KeyBytes: x.KeyBytes,
			ResultBytes: x.ResultBytes,
		}
		candidates.PerByte = append(candidates.PerByte, obx.Candidates())
	}

	return candidates, nil
}

func (x *Candidates) MostLikely(expectedDist map[byte]float64) []byte {
	mostLikelyKey := make([]byte, len(x.PerByte))

	for i, candidates := range x.PerByte {
		// Evalute candidates for the n-th byte.

		for j, _ := range candidates {
			x.PerByte[i][j].SetDiff(expectedDist)
		}

		sort.Slice(x.PerByte[i], func (i, j int) bool {
			return candidates[i].Diff < candidates[j].Diff
		})

		mostLikelyKey[i] = x.PerByte[i][0].B
	}

	return mostLikelyKey
}
