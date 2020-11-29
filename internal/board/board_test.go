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

func TestBoard_Valid(t *testing.T) {
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
		g := [][]Square{
			{Yes, Yes, No, Yes, Yes},
			{No, No, No, Yes, Yes},
			{Yes, No, Yes, No, Yes},
			{No, No, No, Yes, Yes},
			{No, No, Yes, Yes, No},
		}
		b := Board{
			Grid:     g,
			RowHints: rowHints,
			ColHints: colHints,
		}

		isValid := b.Valid()
		require.True(t, isValid)
	})
}

func TestValidLine(t *testing.T) {
	t.Run("SuccessfulOnEitherEnd", func(t *testing.T) {
		hint := []int{4, 4}
		line := []Square{Yes, Yes, Yes, Yes, No, No, Yes, Yes, Yes, Yes}

		require.True(t, validLine(hint, line))
	})
	t.Run("SuccessfulLeftPad", func(t *testing.T) {
		hint := []int{4, 4}
		line := []Square{No, Yes, Yes, Yes, Yes, No, Yes, Yes, Yes, Yes}

		require.True(t, validLine(hint, line))
	})
	t.Run("SuccessfulRightPad", func(t *testing.T) {
		hint := []int{4, 4}
		line := []Square{Yes, Yes, Yes, Yes, No, Yes, Yes, Yes, Yes, No}

		require.True(t, validLine(hint, line))
	})

	t.Run("ShortFirstLength", func(t *testing.T) {
		hint := []int{4, 4}
		line := []Square{Yes, Yes, Yes, No, No, Yes, Yes, Yes, Yes, No}

		require.False(t, validLine(hint, line))
	})
	t.Run("ShortSecondLength", func(t *testing.T) {
		hint := []int{4, 4}
		line := []Square{Yes, Yes, Yes, Yes, No, No, Yes, Yes, Yes, No}

		require.False(t, validLine(hint, line))
	})
	t.Run("ShortFirstLengthAfterPad", func(t *testing.T) {
		hint := []int{4, 4}
		line := []Square{No, Yes, Yes, No, No, Yes, Yes, Yes, Yes, No}

		require.False(t, validLine(hint, line))
	})

	t.Run("LongFirstLength", func(t *testing.T) {
		hint := []int{2, 4}
		line := []Square{Yes, Yes, Yes, No, No, Yes, Yes, Yes, Yes, No}

		require.False(t, validLine(hint, line))
	})
	t.Run("LongSecondLength", func(t *testing.T) {
		hint := []int{2, 3}
		line := []Square{Yes, Yes, No, No, No, Yes, Yes, Yes, Yes, No}

		require.False(t, validLine(hint, line))
	})
	t.Run("LongFirstLengthAfterPad", func(t *testing.T) {
		hint := []int{2, 4}
		line := []Square{No, Yes, Yes, Yes, No, Yes, Yes, Yes, Yes, No}

		require.False(t, validLine(hint, line))
	})

	t.Run("AnotherYesAfterHintsComplete", func(t *testing.T) {
		hint := []int{1, 1}
		line := []Square{Yes, No, Yes, No, Yes}

		require.False(t, validLine(hint, line))
	})

	t.Run("Unknown", func(t *testing.T) {
		hint := []int{1, 1}
		line := []Square{Yes, No, Yes, No, Unknown}

		require.False(t, validLine(hint, line))
	})
}
