package q004

import "strings"

// Answer of 问题4，替换空格
// 使用两个指针遍历一次减少复杂度。
// 1. 遍历一遍字符串计算需要替换的次数；
// 2. 根据需要替换的次数新建或扩容容器；
// 3. 使用两个指针分别指向新旧容器的末尾；
// 4. 将旧容器末尾内容依次写入新容器末尾，遇到需要替换的内容时进行替换；
func Answer(inputStr string) (outputStr string) {
	old := " "
	new := "%20"
	replaceCount := strings.Count(inputStr, " ")
	outputByte := make([]byte, len(inputStr)+replaceCount*(len(new)-len(old)))
	i := len(inputStr) - 1
	j := len(outputByte) - 1
	for i > -1 {
		if inputStr[i] == old[0] {
			copy(outputByte[j-(len(new)-len(old)):j+1], new)
			j -= len(new) - len(old)
		} else {
			outputByte[j] = inputStr[i]
		}
		i--
		j--
	}
	return string(outputByte)
}

