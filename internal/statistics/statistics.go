package statistics

import (
	"math"
)

// RelByteDist returns the relative frequency distribution of the given bytes.
func RelByteDist(bytes []byte) map[byte]float64 {
	total := float64(len(bytes))
	absDist := make(map[byte]int)
	relDist := make(map[byte]float64)

	for _, b := range bytes {
		absDist[b] += 1
	}
	
	for b, c := range absDist {
		relDist[b] = float64(c) / total
	}

	return relDist
}

// ChiSqure calculates the following measure for the similarity of the two given distributions,
// observed and theoretical:
//   sum_(i=0)^n ((observed[i] - theoretical[i])^2 / theoretical[i])
func ChiSquare(observed map[byte]float64, theoretical map[byte]float64) float64 {
	res := 0.0
	for i, t := range theoretical {
		res += math.Pow(observed[i] - t, 2) / t
	}
	return res
}
