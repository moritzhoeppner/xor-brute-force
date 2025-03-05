// Package obxor provides routines for brute-forcing a one-byte XOR cipher.
//
// A one-byte XOR cipher generates a ciphertext by XORing each byte of the plaintext with a fixed
// key byte.
package obxor

import (
	"errors"
	"slices"
	"github.com/moritzhoeppner/xor-brute-force/internal/statistics"
)

type Obxor struct {
	Ciphertext  []byte
	KeyBytes    []byte
	ResultBytes []byte
}

type Candidate struct {
	B      byte
	Result []byte
	Diff   float64
}

func (x *Candidate) SetDiff(expectedDist map[byte]float64) float64 {
	actualDist := statistics.RelByteDist(x.Result)
	x.Diff = statistics.ChiSquare(actualDist, expectedDist)
	return x.Diff
}

func (x *Obxor) Candidates() []Candidate {
	candidates := []Candidate{}

	for _, b := range x.KeyBytes {
		if res, err := x.try(b); err == nil {
			candidates = append(candidates, Candidate{B: b, Result: res})
		}
	}

	return candidates
}

func (x *Obxor) try(candidate byte) ([]byte, error) {
	result := make([]byte, len(x.Ciphertext))

	for i := range result {
		result[i] = candidate ^ x.Ciphertext[i]

		if !slices.Contains(x.ResultBytes, result[i]) {
			return result, errors.New("Result contains not allowed bytes.")
		}
	}

	return result, nil
}
