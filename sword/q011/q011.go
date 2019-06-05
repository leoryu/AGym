package q011

func Answer(base float64, exponent int) (result float64) {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	if exponent == -1 {
		return 1 / base
	}
	result = Answer(base, exponent>>1)
	result *= result
	if (exponent & 1) == 1 {
		result *= base
	}
	return
}

