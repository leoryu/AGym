package main

import (
	"math/big"
)

func FibonacciSequence(n int) (numberN *big.Int) {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}
	num1 := big.NewInt(0)
	num2 := big.NewInt(1)
	for i := 2; i <= n; i++ {
		tmpResult := big.NewInt(0)
		tmpResult.Add(num1, num2)
		num1 = num2
		num2 = tmpResult
	}
	return num2
}

