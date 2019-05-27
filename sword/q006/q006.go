package q006

type BinaryTreeNode struct {
	Value     int
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

// AnswerWithRecursion of 问题6，注意有一个特殊条件是遍历结果中不含重复数字。
// 前序遍历顺序为根左右，中序遍历顺序为左根右。根据前面的不重复条件，我们可以找到中序遍历中的根节点位置，并确定左右树的集合。
// 之后分别在左右树中递归调用函数即可。
func AnswerWithRecursion(preOrder []int, inOrder []int) (root *BinaryTreeNode) {
	if preOrder == nil || len(preOrder) == 0 {
		return
	}
	root = new(BinaryTreeNode)
	root.Value = preOrder[0]
	var rootPosition int
	for i, v := range inOrder {
		if v == root.Value {
			rootPosition = i
		}
	}
	// leftLen是多余的，但有助于理解
	leftLen := rootPosition
	rightLen := len(preOrder) - rootPosition - 1
	if leftLen > 0 {
		root.LeftNode = AnswerWithRecursion(preOrder[1:rootPosition+1], inOrder[:rootPosition])
	}
	if rightLen > 0 {
		root.RightNode = AnswerWithRecursion(preOrder[rootPosition+1:], inOrder[rootPosition+1:])
	}
	return
}

