package tree

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestInsertion(t *testing.T) {
	g := Goblin(t)
	n1 := NewNode(10)

	g.Describe("Node insertion", func() {
		// Passing Test
		g.It("Should insert node", func() {
			g.Assert(n1.Insert(9).Label).Equal(n1.Label)
		})
	})
}
