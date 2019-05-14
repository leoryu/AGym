package main

import "github.com/leoryu/AGym/tree"

func RebuildTree(preOrder []interface{}, inOrder []interface{}) (root *tree.BinaryTree) {
	var rebuildFunc func(pre []interface{}, in []interface{}) (t *tree.BinaryTree)
	rebuildFunc = func(pre []interface{}, in []interface{}) (t *tree.BinaryTree) {
		if len(pre) == 0 {
			return
		}
		t = &tree.BinaryTree{Value: pre[0]}
		if len(pre) == 1 {
			return
		}
		leftLen := 0
		for _, value := range in {
			if value == pre[0] {
				break
			}
			leftLen++
		}
		rightLen := len(in) - leftLen - 1
		if leftLen > 0 {
			t.Left = rebuildFunc(pre[1:leftLen+1], in[:leftLen])
		}
		if rightLen > 0 {
			t.Right = rebuildFunc(pre[leftLen+1:], in[leftLen+1:])
		}
		return
	}
	root = rebuildFunc(preOrder, inOrder)
	return
}

