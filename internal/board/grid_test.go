package board

import "testing"

func TestGrid_Print(t *testing.T) {
	g := New(5, 5)

	g[0][1] = Yes

	g.Print()
}
