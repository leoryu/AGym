package main

import (
	"reflect"
	"testing"

	"github.com/leoryu/AGym/tree"
)

func TestRebuildTree(t *testing.T) {
	type args struct {
		preOrder []interface{}
		inOrder  []interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantRoot *tree.BinaryTree
	}{
		{
			"Test case 1",
			args{
				preOrder: []interface{}{10, 6, 4, 8, 14, 12, 16},
				inOrder: []interface{}{4, 6, 8, 10, 12, 14, 16},
			},
			&treeRoot,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := PostOrderNoRecursion(RebuildTree(tt.args.preOrder, tt.args.inOrder)); !reflect.DeepEqual(gotRoot, PostOrderNoRecursion(tt.wantRoot)) {
				t.Errorf("RebuildTree() = %v, want %v", gotRoot, PostOrderNoRecursion(tt.wantRoot))
			}
		})
	}
}

