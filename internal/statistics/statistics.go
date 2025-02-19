package statistics

import (
	"math"
)

// Returns the relative frequency distribution of the bytes in {byte}.
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

func ChiSquare(observed map[byte]float64, theoretical map[byte]float64) float64 {
	res := 0.0
	for i, t := range theoretical {
		res += math.Pow(observed[i] - t, 2) / t
	}
	return res
}
