package main

import (
	"github.com/leoryu/AGym/stack"
	"github.com/leoryu/AGym/tree"
)

func PreOrderNoRecursion(t *tree.BinaryTree) (values []interface{}) {
	var s stack.Stack
	for t != nil || len(s) != 0 {
		for t != nil {
			values = append(values, t.Value)
			s.Push(t)
			t = t.Left
		}
		if len(s) != 0 {
			t = s.Pop().(*tree.BinaryTree).Right
		}
	}
	return
}

func InOrderNoRecursion(t *tree.BinaryTree) (values []interface{}) {
	var s stack.Stack
	for t != nil || len(s) != 0 {
		for t != nil {
			s.Push(t)
			t = t.Left
		}
		if len(s) != 0 {
			t = s.Pop().(*tree.BinaryTree)
			values = append(values, t.Value)
			t = t.Right
		}
	}
	return
}

func PostOrderNoRecursion(t *tree.BinaryTree) (values []interface{}) {
	var (
		s             stack.Stack
		candidateLeaf *tree.BinaryTree
		preAddedLeaf  *tree.BinaryTree
	)
	for t != nil || len(s) != 0 {
		for t != nil {
			s.Push(t)
			t = t.Left
		}
		candidateLeaf = s.Pop().(*tree.BinaryTree)
		if (candidateLeaf.Left == nil && candidateLeaf.Right == nil) ||
			(candidateLeaf.Right == nil && preAddedLeaf == candidateLeaf.Left) ||
			(preAddedLeaf == candidateLeaf.Right) {
			values = append(values, candidateLeaf.Value)
			preAddedLeaf = candidateLeaf
		} else {
			s.Push(candidateLeaf)
			t = candidateLeaf.Right
		}
	}
	return
}

func PreOrderRecursion(t *tree.BinaryTree) (values []interface{}) {
	var orderFunc func(*tree.BinaryTree)
	orderFunc = func(bt *tree.BinaryTree) {
		values = append(values, bt.Value)
		if bt.Left != nil {
			orderFunc(bt.Left)
		}
		if bt.Right != nil {
			orderFunc(bt.Right)
		}
	}
	orderFunc(t)
	return
}

func InOrderRecursion(t *tree.BinaryTree) (values []interface{}) {
	var orderFunc func(*tree.BinaryTree)
	orderFunc = func(bt *tree.BinaryTree) {
		if bt.Left != nil {
			orderFunc(bt.Left)
		}
		values = append(values, bt.Value)
		if bt.Right != nil {
			orderFunc(bt.Right)
		}
	}
	orderFunc(t)
	return
}

func PostOrderRecursion(t *tree.BinaryTree) (values []interface{}) {
	var orderFunc func(*tree.BinaryTree)
	orderFunc = func(bt *tree.BinaryTree) {
		if bt.Left != nil {
			orderFunc(bt.Left)
		}
		if bt.Right != nil {
			orderFunc(bt.Right)
		}
		values = append(values, bt.Value)
	}
	orderFunc(t)
	return
}

