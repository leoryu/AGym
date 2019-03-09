package main

import (
	"fmt"

	"github.com/leoryu/AGym/linkedlist"
	"github.com/leoryu/AGym/stack"
)

func RecersePrint(head *linkedlist.Single) (result string) {
	i := head
	var s stack.Stack
	var printResult []byte
	for i != nil {
		s.Push(i)
		i = i.Next
	}
	j := s.Pop()
	for j != nil {
		printResult = append(printResult, fmt.Sprintln(j.(*linkedlist.Single).Data)[:]...)
		j = s.Pop()
	}
	return string(printResult)
}

