package q009

import (
	"math/big"
)

func Answer(n int) (numberN *big.Int) {
	num1 := big.NewInt(0)
	num2 := big.NewInt(1)
	if n <= 0 {
		return num1
	}
	if n == 1 {
		return num2
	}
	for i := 2; i <= n; i++ {
		tmpResult := big.NewInt(0)
		tmpResult.Add(num1, num2)
		num1 = num2
		num2 = tmpResult
	}
	return num2
}

