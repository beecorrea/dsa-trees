package main

import (
	"fmt"

	"github.com/beecorrea/bst/tree"
)

func main() {
	n1 := tree.NewNode(5)
	n2 := tree.NewNode(2)
	n3 := tree.NewNode(1)
	// n4 := tree.NewNode(6)

	n1.Left = n2
	n2.Left = n3
	// n1.Right = n4

	n1.Preorder()
	fmt.Println()

	fmt.Println(n1.CalcHeight())
	fmt.Println(n1.CalcBalance())
	root := n1.RightRotation()

	fmt.Println()
	root.Preorder()
}
