package q018

type BinaryTreeNode struct {
	Value     int
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

func Answer(a, b *BinaryTreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Value == b.Value {
		return isSub(a, b)
	}
	return Answer(a.LeftNode, b) || Answer(a.RightNode, b)
}

func isSub(a, b *BinaryTreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if b.Value != a.Value {
		return false
	}
	return isSub(a.LeftNode, b.LeftNode) && isSub(a.RightNode, b.RightNode)
}
