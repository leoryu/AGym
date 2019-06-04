package q010

func Answer(checkedNum int) (numOfOne int) {
	if checkedNum < 0 {
		checkedNum = 0 - checkedNum
	}
	for checkedNum != 0 {
		checkedNum = checkedNum & (checkedNum - 1)
		numOfOne++
	}
	return
}

