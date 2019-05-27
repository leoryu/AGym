package q006

import (
	"reflect"
	"testing"
)

var (
	node1 *BinaryTreeNode
	node2 *BinaryTreeNode
)

func init() {
	node1 = new(BinaryTreeNode)
	node1.Value = 1
	node1.LeftNode = new(BinaryTreeNode)
	node1.LeftNode.Value = 2
	node1.LeftNode.LeftNode = new(BinaryTreeNode)
	node1.LeftNode.LeftNode.Value = 4
	node1.LeftNode.LeftNode.RightNode = new(BinaryTreeNode)
	node1.LeftNode.LeftNode.RightNode.Value = 7
	node1.RightNode = new(BinaryTreeNode)
	node1.RightNode.Value = 3
	node1.RightNode.LeftNode = new(BinaryTreeNode)
	node1.RightNode.LeftNode.Value = 5
	node1.RightNode.RightNode = new(BinaryTreeNode)
	node1.RightNode.RightNode.Value = 6
	node1.RightNode.RightNode.LeftNode = new(BinaryTreeNode)
	node1.RightNode.RightNode.LeftNode.Value = 8
}

func TestAnswerWithRecursion(t *testing.T) {
	type args struct {
		preOrder []int
		inOrder  []int
	}
	tests := []struct {
		name     string
		args     args
		wantRoot *BinaryTreeNode
	}{
		{
			"Function test 1",
			args{
				preOrder: []int{1, 2, 4, 7, 3, 5, 6, 8},
				inOrder:  []int{4, 7, 2, 1, 5, 3, 8, 6},
			},
			node1,
		},
		{
			"Empty",
			args{
				preOrder: []int{},
				inOrder:  []int{},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := AnswerWithRecursion(tt.args.preOrder, tt.args.inOrder); !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("AnswerWithRecursion() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
	}
}

