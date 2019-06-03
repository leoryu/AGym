package q008

func Answer(input []int) (output []int) {
	if input == nil || len(input) == 0 {
		return []int{0}
	}
	if len(input) == 1 {
		return input
	}
	lastPoint := len(input) - 1
	middlePoint := len(input) / 2
	if input[lastPoint] > input[0] {
		return input[:1]
	}
	if input[0] == input[middlePoint] && input[0] == input[lastPoint] {
		return Answer(input[1:])
	}
	if input[middlePoint] >= input[0] {
		return Answer(input[middlePoint+1:])
	}
	return Answer(input[1 : middlePoint+1])
}
