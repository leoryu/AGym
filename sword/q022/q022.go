package q022

type Stack []int

func (s Stack) Top() *int {
	if len(s) == 0 {
		return nil
	}
	return &s[len(s)-1]
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *int {
	oldStack := *s
	if len(*s) == 0 {
		return nil
	}
	*s = oldStack[:len(oldStack)-1]
	return &oldStack[len(oldStack)-1]
}

func Answer(in, out []int) bool {
	if len(in) != len(out) {
		return false
	}
	var (
		i, j  int
		stack Stack
	)
	for i < len(in) {
		stack.Push(in[i])
		i++
		if *stack.Top() != out[j] {
			continue
		}
		for j < len(out) {
			if stack.Top() == nil || *stack.Top() != out[j] {
				break
			}
			stack.Pop()
			j++
		}
	}

	if stack.Top() == nil {
		return true
	}
	return false
}
