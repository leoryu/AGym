package q005

type NodeList struct {
	Value    int
	NextNode *NodeList
}

// AnswerWithRecursion of 问题5，使用递归，从头节点开始，以下一个节点为入参递归调用，直到入参为nil为止。
func AnswerWithRecursion(head *NodeList) (result []int) {
	if head == nil {
		return []int{}
	}
	nextResult := AnswerWithRecursion(head.NextNode)
	if len(nextResult) != 0 {
		return append(nextResult, head.Value)
	}
	return []int{head.Value}
}

// AnswerWithoutRecursion of 问题5，未使用递归。
func AnswerWithoutRecursion(head *NodeList) (result []int) {
	count := 0
	node := head
	for node != nil {
		count++
		node = node.NextNode
	}
	result = make([]int, count)
	for head != nil {
		result[count-1] = head.Value
		head = head.NextNode
		count--
	}
	return
}

