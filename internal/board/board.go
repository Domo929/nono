package board

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	// ErrRowMismatch is an error for when the number of rows and row hints are not the same
	ErrRowMismatch = errors.New("mismatch between number of rows and row hints")

	// ErrColMismatch is an error for when the number of cols and col hints are not the same
	ErrColMismatch = errors.New("mismatch between number of cols and col hints")
)

// Board is a struct that contains the hints for each row and column, as well as the current state of the board
type Board struct {
	Grid     Grid    `json:"-"`
	RowHints [][]int `json:"row_hints"`
	ColHints [][]int `json:"col_hints"`
}

// New takes in the x and y size and the hints and returns a board
// an error is returned if the length of the hints does not match the grid length
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

// NewFromFile opens a board game from a json file and returns it
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

// Solved checks a board to see if it's solved
func (b Board) Solved() bool {
	for rowNdx, row := range b.Grid {
		if !validLine(b.RowHints[rowNdx], row) {
			return false
		}
	}
	for i := 0; i < len(b.Grid[0]); i++ {
		if !validLine(b.ColHints[i], getCol(b, i)) {
			return false
		}
	}

	return true
}

// getCol extracts a column from the game board as a singular slice for validation
func getCol(b Board, colNdx int) []Square {
	col := make([]Square, 0)
	for _, row := range b.Grid {
		col = append(col, row[colNdx])
	}
	return col
}

// validLine checks a line to see if the layout matches the provided hints
func validLine(hint []int, line []Square) bool {
	if len(hint) == 0 || len(line) == 0 {
		return false
	}
	hintNdx := 0
	curLength := 0
	curHint := &hint[hintNdx]
	afterYes := false
	for _, sq := range line {
		switch sq {
		case Unknown:
			return false
		case Yes:
			curLength++
			afterYes = true
			if curHint == nil {
				return false
			}
			if curLength == *curHint {
				curHint, hintNdx = nextHint(hint, hintNdx)
			}
		case No:
			if curHint != nil && curLength != *curHint && afterYes {
				return false
			}
			curLength = 0
			afterYes = false
		}
	}
	return true
}

// nextHint is a simple helper that checks whether there is a next hint
func nextHint(hint []int, curNdx int) (*int, int) {
	nextNdx := curNdx + 1
	if nextNdx >= len(hint) {
		return nil, curNdx
	}
	return &hint[nextNdx], nextNdx
}
