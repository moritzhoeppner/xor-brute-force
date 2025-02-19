package mbxor

import (
	"github.com/moritzhoeppner/xor-brute-force/internal/utils"
	"github.com/moritzhoeppner/xor-brute-force/internal/obxor"
)

type Mbxor struct {
	Ciphertexts [][]byte
	KeyBytes    []byte
	ResultBytes []byte
}

func (x *Mbxor) Candidates() [][]obxor.Candidate {
	candidates := [][]obxor.Candidate{}

	// Transpose Ciphertexts. The elements of the result have the same one-byte key.
	obCiphertexts := utils.Transpose(x.Ciphertexts)

	for _, ciphertext := range obCiphertexts {
		obx := obxor.Obxor{
			Ciphertext: ciphertext,
			KeyBytes: x.KeyBytes,
			ResultBytes: x.ResultBytes,
		}
		candidates = append(candidates, obx.Candidates())
	}

	return candidates
}
