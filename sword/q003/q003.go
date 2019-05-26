package q003

// Answer of 剑值Offer问题3，二维数组中的查找
// 采用分治思想。从右上角开始：
// 1. 如果等于我们找的数字，返回true；
// 2. 如果小于我们找的数字，该行可以删除，在新数组中继续找；
// 3. 如果大于我们找的数字，该列可以删除，在新数组中继续找；
// 4. 当数组中没有数字时，返回false；
func Answer(inputArray [][]int, expectedNumber int) (isExisted bool) {
	if inputArray == nil || len(inputArray) == 0 {
		return false
	}
	h := len(inputArray)
	w := len(inputArray[0])
	for i, j := 0, w-1; i < h && j >= 0; {
		if inputArray[i][j] == expectedNumber {
			return true
		}
		if inputArray[i][j] < expectedNumber {
			i++
		}
		if inputArray[i][j] > expectedNumber {
			j--
		}
	}
	return
}
