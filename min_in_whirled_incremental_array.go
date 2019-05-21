package main

func MinInWhirledIncrementalArray(incrementalArray []int) (minNumber int) {
	var binarySearchFunc func([]int) []int
	binarySearchFunc = func(inputArray []int) (outpuArray []int) {
		arrayLen := len(inputArray)
		lastIndex := arrayLen - 1
		middleIndex := arrayLen / 2
		if len(inputArray) == 1 {
			return inputArray
		}
		if inputArray[lastIndex] > inputArray[0] {
			return inputArray[:1]
		}
		if inputArray[0] == inputArray[lastIndex] && inputArray[0] == inputArray[middleIndex] {
			return binarySearchFunc(inputArray[1:])
		}
		if inputArray[middleIndex] >= inputArray[0] {
			return binarySearchFunc(inputArray[middleIndex+1:])
		}
		return binarySearchFunc(inputArray[:middleIndex+1])
	}
	return binarySearchFunc(incrementalArray)[0]
}

