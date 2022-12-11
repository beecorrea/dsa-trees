package tree

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("BST Search", func() {
		tr := leftDiagonal()
		// Passing Test
		g.It("Should find node 2", func() {
			g.Assert(tr.Search(2).Label).Equal(2)
		})

		// Failing Test
		g.It("Should not find node 7", func() {
			g.Assert(tr.Search(7)).IsNil()
		})
	})
}

func leftDiagonal() *Node {
	n1 := NewNode(5)
	n2 := NewNode(2)
	n3 := NewNode(1)

	n1.Left = n2
	n2.Left = n3

	return n1
}
