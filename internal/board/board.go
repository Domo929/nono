package board

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrRowMismatch = errors.New("mismatch between number of rows and row hints")
	ErrColMismatch = errors.New("mismatch between number of cols and col hints")
)

type Board struct {
	Grid     Grid    `json:"-"`
	RowHints [][]int `json:"row_hints"`
	ColHints [][]int `json:"col_hints"`
}

func New(x, y int, rowHints, colHints [][]int) (*Board, error) {
	if len(rowHints) != y {
		return nil, ErrRowMismatch
	}
	if len(colHints) != x {
		return nil, ErrColMismatch
	}
	return &Board{
		Grid:     newGrid(x, y),
		RowHints: rowHints,
		ColHints: colHints,
	}, nil
}

func NewFromFile(path string) (*Board, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b := new(Board)
	if err := json.NewDecoder(f).Decode(b); err != nil {
		return nil, err
	}
	b.Grid = newGrid(len(b.RowHints), len(b.ColHints))

	return b, nil
}
