package main

func CustomPower(base float64, exponent int) (result float64) {
	if exponent == 0 {
		return 1
	}
	if exponent == 1{
		return base
	}
	result = CustomPower(base, exponent >> 1)
	result *= result
	if (exponent & 1) == 1 {
		result *= base
	}
	return
}

