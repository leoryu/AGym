package q016

type NodeList struct {
	Value    int
	NextNode *NodeList
}

func Answer(head *NodeList) (target *NodeList) {
	if head == nil || head.NextNode == nil {
		return head
	}
	var tmpNode *NodeList
	for head != nil {
		next := head.NextNode
		head.NextNode = tmpNode
		tmpNode = head
		head = next
	}
	return tmpNode
}
