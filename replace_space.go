package main

import "strings"

func ReplaceSpace(str string) (newStr string) {
	old := " "
	new := "%20"
	n := strings.Count(str, " ")
	t := make([]byte, len(str)+n*(len(new)-len(old)))
	i := len(str) - 1
	j := len(t) - 1
	for i > -1 {
		if str[i] == old[0] {
			copy(t[j-(len(new)-len(old)):j+1], new)
			j -= len(new) - len(old)
		} else {
			t[j] = str[i]
		}
		i--
		j--
	}
	return string(t)
}

