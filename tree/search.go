package tree

func (n *Node) Search(label int) *Node {
	if n == nil {
		return nil
	}

	if label == n.Label {
		return n
	}

	if n.Left != nil && label < n.Label {
		return n.Left.Search(label)
	}

	if n.Right != nil && label > n.Label {
		return n.Right.Search(label)
	}

	return nil
}
