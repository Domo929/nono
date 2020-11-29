package board

import (
	"fmt"
	"io"
	"os"
)

// Square identifies what value a grid square holds
type Square uint8

const (
	Unknown Square = iota
	Yes
	No
)
const (
	unknownChar = "~"
	yesChar     = "#"
	noChar      = "X"
)

// Grid is a slice of slices denoting the layout of the board
// The axes are [y][x]Square
type Grid [][]Square

// newGrid makes the internal slices for a Grid
func newGrid(x, y int) Grid {
	g := make([][]Square, x)
	for i := 0; i < x; i++ {
		g[i] = make([]Square, y)
	}
	return g
}

// Print prints out a simple representation of the board to stdout
// the writer is used to redirect to a different writer than stdout. If w is nil,
// stdout is used
func (g Grid) Print(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}
	gridStr := ""
	for _, row := range g {
		for _, sq := range row {
			var char string
			switch sq {
			case Unknown:
				char = unknownChar
			case Yes:
				char = yesChar
			case No:
				char = noChar
			}
			gridStr += char
		}
		gridStr += "\n"
	}
	fmt.Fprint(w, gridStr)
}
