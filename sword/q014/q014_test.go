package q014

import (
	"reflect"
	"testing"
)

func TestAnswerTwoPointer(t *testing.T) {
	type args struct {
		inputList []int
	}
	tests := []struct {
		name           string
		args           args
		wantOutputList []int
	}{
		{
			"Function test 1",
			args{
				[]int{1, 2, 3, 4, 5},
			},
			[]int{1, 5, 3, 4, 2},
		},
		{
			"0 test",
			args{
				[]int{},
			},
			[]int{},
		},
		{
			"Even test",
			args{
				[]int{2, 4, 6, 8},
			},
			[]int{2, 4, 6, 8},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputList := AnswerTwoPointer(tt.args.inputList); !reflect.DeepEqual(gotOutputList, tt.wantOutputList) {
				t.Errorf("AnswerTwoPointer() = %v, want %v", gotOutputList, tt.wantOutputList)
			}
		})
	}
}

func TestAnswerBubble(t *testing.T) {
	type args struct {
		inputList []int
	}
	tests := []struct {
		name           string
		args           args
		wantOutputList []int
	}{
		{
			"Function test 1",
			args{
				[]int{1, 2, 3, 4, 5},
			},
			[]int{1, 3, 5, 2, 4},
		},
		{
			"Function test 1",
			args{
				[]int{2, 3, 4},
			},
			[]int{3, 2, 4},
		},
		{
			"0 test",
			args{
				[]int{},
			},
			[]int{},
		},
		{
			"Even test",
			args{
				[]int{2, 4, 6, 8},
			},
			[]int{2, 4, 6, 8},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputList := AnswerBubble(tt.args.inputList); !reflect.DeepEqual(gotOutputList, tt.wantOutputList) {
				t.Errorf("AnswerBubble() = %v, want %v", gotOutputList, tt.wantOutputList)
			}
		})
	}
}
