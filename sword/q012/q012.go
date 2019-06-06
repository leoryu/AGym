package q012

import (
	"strconv"
)

var results []string

func Answer(n int) []string {
	results = []string{}
	if n <= 0 {
		return results
	}
	for i := 0; i < 10; i++ {
		AddResult(n, strconv.Itoa(i))
	}
	return results[1:]
}

func AddResult(length int, tmpResult string) {
	if len(tmpResult) == length {
		results = append(results, CleanResult(tmpResult))
		return
	}
	for i := 0; i < 10; i++ {
		AddResult(length, tmpResult+strconv.Itoa(i))
	}
}

func CleanResult(input string) (output string) {
	output = input
	for _, v := range []byte(input) {
		if v != []byte("0")[0] {
			break
		}
		output = output[1:]
	}
	return
}

