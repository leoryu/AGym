package q019

type BinaryTreeNode struct {
	Value     int
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

func Answer(treeNode *BinaryTreeNode) (mirrorTreeNode *BinaryTreeNode) {
	mirror(treeNode)
	return treeNode
}

func mirror(treeNode *BinaryTreeNode) {
	if treeNode == nil {
		return
	}
	treeNode.LeftNode, treeNode.RightNode = treeNode.RightNode, treeNode.LeftNode
	mirror(treeNode.LeftNode)
	mirror(treeNode.RightNode)
}
