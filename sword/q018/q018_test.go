package q018

import "testing"

func TestAnswer(t *testing.T) {
	tree1 := &BinaryTreeNode{Value: 1}
	tree2 := &BinaryTreeNode{Value: 2}
	tree3 := &BinaryTreeNode{Value: 3}
	tree4 := &BinaryTreeNode{Value: 4}
	tree5 := &BinaryTreeNode{Value: 5}
	tree4.LeftNode = tree5
	tree3.RightNode = tree4
	tree2.LeftNode = tree4
	tree2.RightNode = tree3
	tree1.RightNode = tree5
	tree1.LeftNode = tree2
	type args struct {
		a *BinaryTreeNode
		b *BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Function test 1",
			args{
				a: tree1,
				b: tree2,
			},
			true,
		},
		{
			"Function test 2",
			args{
				a: tree2,
				b: tree1,
			},
			false,
		},
		{
			"Function test 3",
			args{
				a: tree1,
				b: tree1,
			},
			true,
		},
		{
			"Empty test 1",
			args{
				a: nil,
				b: nil,
			},
			true,
		},
		{
			"Empty test 2",
			args{
				a: tree5,
				b: nil,
			},
			true,
		},
		{
			"Empty test 3",
			args{
				a: nil,
				b: tree5,
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Answer(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Answer() = %v, want %v", got, tt.want)
			}
		})
	}
}
