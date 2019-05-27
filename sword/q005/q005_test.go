package q005

import (
	"reflect"
	"testing"
)

func TestAnswerWithRecursion(t *testing.T) {
	type args struct {
		head *NodeList
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			"Function test 1",
			args{
				head: &NodeList{
					Value: 1,
					NextNode: &NodeList{
						Value: 2,
						NextNode: &NodeList{
							Value:    3,
							NextNode: nil,
						},
					},
				},
			},
			[]int{3, 2, 1},
		},
		{
			"Function test 2",
			args{
				head: &NodeList{
					Value:    1,
					NextNode: nil,
				},
			},
			[]int{1},
		},
		{
			"Empty",
			args{
				head: nil,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AnswerWithRecursion(tt.args.head); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AnswerWithRecursion() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestAnswerWithoutRecursion(t *testing.T) {
	type args struct {
		head *NodeList
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			"Function test 1",
			args{
				head: &NodeList{
					Value: 1,
					NextNode: &NodeList{
						Value: 2,
						NextNode: &NodeList{
							Value:    3,
							NextNode: nil,
						},
					},
				},
			},
			[]int{3, 2, 1},
		},
		{
			"Function test 2",
			args{
				head: &NodeList{
					Value:    1,
					NextNode: nil,
				},
			},
			[]int{1},
		},
		{
			"Empty",
			args{
				head: nil,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AnswerWithoutRecursion(tt.args.head); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AnswerWithoutRecursion() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

