package stack

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (item interface{}) {
	self := *s
	if len(self) == 0 {
		return
	}
	item = self[len(self)-1]
	*s = self[:len(self)-1]
	return
}

