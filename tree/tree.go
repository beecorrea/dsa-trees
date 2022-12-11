package tree

import (
	"fmt"
	"math"
)

type Node struct {
	Label int
	Left  *Node
	Right *Node

	Height  int
	Balance int
}

func NewNode(label int) *Node {
	return &Node{label, nil, nil, 0, 0}
}

func (n *Node) CalcBalance() int {
	bl := int(float64(n.Left.CalcHeight()) - float64(n.Right.CalcHeight()))
	n.Balance = bl
	return bl
}

func (n *Node) IsUnbalanced() bool {
	return n.Balance > 1 || n.Balance < -1
}

func (n *Node) CalcHeight() int {
	if n == nil {
		return -1
	}

	leftH := n.Left.CalcHeight()
	rightH := n.Right.CalcHeight()

	h := 1 + int(math.Max(float64(leftH), float64(rightH)))
	n.Height = h
	return h
}

// Root -> Left -> Right
func (n *Node) Preorder() {
	if n != nil {
		fmt.Printf("%d ", n.Label)
	}

	if n.Left != nil {
		n.Left.Preorder()
	}

	if n.Right != nil {
		n.Right.Preorder()
	}
}
