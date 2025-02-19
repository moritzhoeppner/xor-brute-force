package obxor

import (
	"errors"
	"slices"
)

type Obxor struct {
	Ciphertext  []byte
	KeyBytes    []byte
	ResultBytes []byte
}

type Candidate struct {
	B      byte
	Result []byte
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
