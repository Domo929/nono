package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("Successful", func(t *testing.T) {
		rowHints := [][]int{
			{2, 2},
			{2},
			{1, 1, 1},
			{2},
			{2},
		}
		colHints := [][]int{
			{1, 1},
			{1},
			{1, 1},
			{2, 2},
			{4},
		}
		b, err := New(5, 5, rowHints, colHints)
		require.NoError(t, err)
		assert.Len(t, b.ColHints, 5)
		assert.Len(t, b.RowHints, 5)
		assert.Len(t, b.Grid, 5)
		assert.Len(t, b.Grid[:], 5)
	})

	t.Run("ColumnMismatch", func(t *testing.T) {
		rowHints := [][]int{
			{2, 2},
			{2},
			{1, 1, 1},
			{2},
			{2},
		}
		colHints := [][]int{
			{1, 1},
			{1},
			{1, 1},
			{2, 2},
		}
		b, err := New(5, 5, rowHints, colHints)
		require.Equal(t, ErrColMismatch, err)
		require.Nil(t, b)
	})

	t.Run("RowMismatch", func(t *testing.T) {
		rowHints := [][]int{
			{2, 2},
			{2},
			{1, 1, 1},
			{2},
		}
		colHints := [][]int{
			{1, 1},
			{1},
			{1, 1},
			{2, 2},
			{4},
		}
		b, err := New(5, 5, rowHints, colHints)
		require.Equal(t, ErrRowMismatch, err)
		require.Nil(t, b)
	})

}

func TestNewFromFile(t *testing.T) {
	b, err := NewFromFile("../../examples/5x5.json")
	require.NoError(t, err)
	assert.Len(t, b.ColHints, 5)
	assert.Len(t, b.RowHints, 5)
	assert.Len(t, b.Grid, 5)
	assert.Len(t, b.Grid[:], 5)
}
