package utils

import (
	"errors"
)

// Transpose switches rows and columns of the given two-dimensional slice. If such a slice is
// interpreted as matrix, the result is the transpose matrix.
func Transpose[T any](s [][]T) ([][]T, error) {
	width := len(s)
	
	if width == 0 {
		return [][]T{}, nil
	}

	height := len(s[0])

	t := make([][]T, height)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if len(s[j]) != height {
				return nil, errors.New("All elements must have the same length.")
			}
			t[i] = append(t[i], s[j][i])
		}
	}

	return t, nil
}
