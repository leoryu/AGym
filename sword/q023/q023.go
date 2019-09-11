package q023

type BinaryTreeNode struct {
	Value     int
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

func Answer(root *BinaryTreeNode) (order []int) {
	if root == nil {
		return
	}
	var nodes []*BinaryTreeNode
	nodes = append(nodes, root)
	for len(nodes) != 0 {
		order = append(order, nodes[0].Value)
		if nodes[0].LeftNode != nil {
			nodes = append(nodes, nodes[0].LeftNode)
		}
		if nodes[0].RightNode != nil {
			nodes = append(nodes, nodes[0].RightNode)
		}
		nodes = nodes[1:]
	}
	return
}
