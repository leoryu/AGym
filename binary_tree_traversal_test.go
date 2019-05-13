package main

import (
	"reflect"
	"testing"

	"github.com/leoryu/AGym/tree"
)

var (
	treeRoot = tree.BinaryTree{
		Value: 10,
	}

	leaf1 = tree.BinaryTree{
		Value: 6,
	}
	leaf2 = tree.BinaryTree{
		Value: 14,
	}
	leaf3 = tree.BinaryTree{
		Value: 4,
	}
	leaf4 = tree.BinaryTree{
		Value: 8,
	}
	leaf5 = tree.BinaryTree{
		Value: 12,
	}
	leaf6 = tree.BinaryTree{
		Value: 16,
	}
)

func init() {
	treeRoot.Left = &leaf1
	treeRoot.Right = &leaf2
	leaf1.Left = &leaf3
	leaf1.Right = &leaf4
	leaf2.Left = &leaf5
	leaf2.Right = &leaf6
}

func TestPreOrderNoRecursion(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{10, 6, 4, 8, 14, 12, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := PreOrderNoRecursion(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PreOrderNoRecursion() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestInOrderNoRecursion(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{4, 6, 8, 10, 12, 14, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := InOrderNoRecursion(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("InOrderNoRecursion() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestPostOrderNoRecursion(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{4, 8, 6, 12, 16, 14, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := PostOrderNoRecursion(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PostOrderNoRecursion() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestPreOrderRecursion(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{10, 6, 4, 8, 14, 12, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := PreOrderRecursion(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PreOrderRecursion() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestInOrderRecursion(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{4, 6, 8, 10, 12, 14, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := InOrderRecursion(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("InOrderRecursion() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestPostOrderRecursion(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{4, 8, 6, 12, 16, 14, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := PostOrderRecursion(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PostOrderRecursion() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestBreadthFirst(t *testing.T) {
	type args struct {
		t *tree.BinaryTree
	}
	tests := []struct {
		name       string
		args       args
		wantValues []interface{}
	}{
		{
			"Test case 1",
			args{
				t: &treeRoot,
			},
			[]interface{}{10, 6, 14, 4, 8, 12, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := BreadthFirst(tt.args.t); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("BreadthFirst() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

