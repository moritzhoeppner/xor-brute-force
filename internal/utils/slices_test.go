package utils

import (
	"testing"
	"gotest.tools/v3/assert"
)

func TestTransposeOneRow(t *testing.T) {
	res, err := Transpose([][]byte{{120, 2}})
	assert.DeepEqual(t, res, [][]byte{{120},{2}})
	assert.Equal(t, err, nil)
}

func TestTransposeOneColumn(t *testing.T) {
	res, err := Transpose([][]byte{{120},{2}})
	assert.DeepEqual(t, res, [][]byte{{120,2}})
	assert.Equal(t, err, nil)
}

func TestTransposeSquare(t *testing.T) {
	res, err := Transpose([][]byte{{1, 2},{3, 4}})
	assert.DeepEqual(t, res, [][]byte{{1,3},{2,4}})
	assert.Equal(t, err, nil)
}

func TestTransposeEmpty(t *testing.T) {
	res, err := Transpose([][]byte{})
	assert.DeepEqual(t, res, [][]byte{})
	assert.Equal(t, err, nil)
}

func TestTransposeWrongSize(t *testing.T) {
	_, err := Transpose([][]byte{{1,2},{1}})
	assert.Error(t, err, "All elements must have the same length.")
}
