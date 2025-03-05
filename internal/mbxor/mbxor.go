// Package mbxor provides routines for brute-forcing a multi-byte XOR cipher.
//
// A multi-byte XOR cipher generates a ciphertext by XORing it bitwise with a fixed key that has a
// byte length > 1.
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

		if len(x.PerByte[i]) > 0 {
			mostLikelyKey[i] = x.PerByte[i][0].B
		}
	}

	return mostLikelyKey
}

func MostLikelyKey(ciphertexts [][]byte, expectedDist map[byte]float64) ([]byte, error) {
	expectedBytes := make([]byte, len(expectedDist))
	i := 0
	for k, _ := range expectedDist {
		expectedBytes[i] = k
		i++
	}

	xor := Mbxor{
		Ciphertexts: ciphertexts,
		KeyBytes: expectedBytes,
		ResultBytes: expectedBytes,
	}

	candidates, err := xor.Candidates()
	if err != nil {
		return nil, err
	}

	return candidates.MostLikely(expectedDist), nil
}
