package q014

func AnswerBubble(inputList []int) (outputList []int) {
	head := 0
	tail := len(inputList) - 1
	for head <= tail {
		for inputList[head]%2 == 1 && head <= tail {
			head++
		}
		for i, j := head, head+1; j < len(inputList) && head <= tail; {
			inputList[i], inputList[j] = inputList[j], inputList[i]
			i++
			j++
		}
		tail--
	}
	return inputList
}

func AnswerTwoPointer(inputList []int) (outputList []int) {
	head := 1
	tail := len(inputList) - 1
	for head < tail {
		for inputList[head]%2 == 1 && head < tail {
			head++
		}
		for inputList[tail]%2 == 0 && head < tail {
			tail--
		}
		inputList[head], inputList[tail] = inputList[tail], inputList[head]
	}
	return inputList
}
