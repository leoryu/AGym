package q017

type NodeList struct {
	Value    int
	NextNode *NodeList
}

func Answer(l1, l2 *NodeList) (ml *NodeList) {
	tmpNode := &NodeList{}
	beforHead := tmpNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmpNode.NextNode = l2
			l2 = l2.NextNode
			break
		}
		if l2 == nil {
			tmpNode.NextNode = l1
			l1 = l1.NextNode
			break
		}
		if l1.Value > l2.Value {
			tmpNode.NextNode = l2
			l2 = l2.NextNode
		} else {
			tmpNode.NextNode = l1
			l1 = l1.NextNode
		}
		tmpNode = tmpNode.NextNode
	}
	return beforHead.NextNode
}

