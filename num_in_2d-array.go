package main

func IsNumIn2dArray(array [][]int, num int) (exist bool) {
	row := len(array) - 1
	if row <= -1 {
		return false
	}
	col := len(array[0]) - 1
	r := row
	c := col
	for r != -1 && c != -1 {
		n := array[row-r][c]
		if num == n {
			return true
		}
		if num > n {
			r--
			continue
		}
		if num < n {
			c--
			continue
		}
	}
	return false
}
