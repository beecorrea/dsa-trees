package tree

func (root *Node) Insert(label int) *Node {
	if root == nil {
		return NewNode(label)
	}

	if root.Label == label {
		return root
	}

	if label < root.Label {
		root.Left = root.Left.Insert(label)
	} else {
		root.Right = root.Right.Insert(label)
	}

	balanceFactor := root.CalcBalance()

	if balanceFactor > 1 {
		leftBF := root.Left.CalcBalance()

		if leftBF == 1 {
			root = root.RightRotation()
		} else {
			root = root.RightLeftRotation()
		}
	} else if balanceFactor < -1 {
		rightBf := root.Right.CalcBalance()
		if rightBf == -1 {
			root = root.LeftRotation()
		} else {
			root.LeftRightRotation()
		}
	}
	return root
}
