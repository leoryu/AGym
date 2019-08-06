package q015

type NodeList struct {
	Value    int
	NextNode *NodeList
}

func Answer(head *NodeList, k int) (target *NodeList) {
	if head == nil || k < 1 {
		return
	}
	target = head
	for k > 1 && head != nil {
		head = head.NextNode
		k--
	}
	if head == nil {
		return nil
	}
	for head.NextNode != nil {
		head = head.NextNode
		target = target.NextNode
	}
	return target
}

