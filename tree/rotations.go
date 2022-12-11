package tree

// Source https://www.scaler.com/topics/data-structures/avl-tree/

// Heavy left sub-tree (BF > 1), clockwise.
func (n *Node) RightRotation() *Node {
	newParent := n.Left
	n.Left = newParent.Right
	newParent.Right = n

	n.CalcHeight()
	n.CalcBalance()
	newParent.CalcHeight()
	newParent.CalcBalance()

	return newParent
}

// Heavy right sub-tree (BF < -1), counter-clockwise
func (n *Node) LeftRotation() *Node {
	newParent := n.Right
	n.Right = newParent.Left
	newParent.Left = n

	n.CalcHeight()
	newParent.CalcHeight()

	n.CalcBalance()
	newParent.CalcBalance()

	return newParent
}

// Transform into left-heavy with left rotation, then apply right rotation.
// Counter-clockwise -> clockwise.
func (n *Node) LeftRightRotation() *Node {
	n.Left = n.Left.LeftRotation()
	newParent := n.RightRotation()
	return newParent
}

// Transform into right-heavy with right rotation, then apply left rotation.
// Clockwise -> Counter-clockwise.
func (n *Node) RightLeftRotation() *Node {
	n.Right = n.Left.RightRotation()
	newParent := n.LeftRotation()
	return newParent
}
