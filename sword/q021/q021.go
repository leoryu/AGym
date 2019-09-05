package q021

type Stack struct {
	stack []item
}

type item struct {
	value, min int
}

func (s *Stack) Push(v int) {
	min := v
	if len(s.stack) > 0 && *s.Min() < min {
		min = *s.Min()
	}
	s.stack = append(s.stack, item{value: v, min: min})
}

func (s *Stack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *Stack) Min() (min *int) {
	if len(s.stack) > 0 {
		return &s.stack[len(s.stack)-1].min
	}
	return
}

func (s *Stack) Top() (top *int) {
	if len(s.stack) > 0 {
		return &s.stack[len(s.stack)-1].value
	}
	return
}
