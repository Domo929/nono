package board

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrid_Print(t *testing.T) {
	g := New(5, 5)

	g[0][1] = Yes

	b := new(bytes.Buffer)
	g.Print(b)

	exp := "~#~~~\n~~~~~\n~~~~~\n~~~~~\n~~~~~\n"
	require.Equal(t, exp, b.String())
}
