package q023

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tree1 := &BinaryTreeNode{Value: 1}
	tree2 := &BinaryTreeNode{Value: 2}
	tree3 := &BinaryTreeNode{Value: 3}
	tree4 := &BinaryTreeNode{Value: 4}
	tree5 := &BinaryTreeNode{Value: 5}
	tree1.LeftNode = tree2
	tree1.RightNode = tree3
	tree2.RightNode = tree4
	tree4.LeftNode = tree5
	type args struct {
		root *BinaryTreeNode
	}
	tests := []struct {
		name      string
		args      args
		wantOrder []int
	}{
		{
			"Function test 1",
			args{root: tree1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"Empty test 1",
			args{root: nil},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrder := Answer(tt.args.root); !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("Answer() = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
