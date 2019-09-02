package q020

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			"Function test 1",
			args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"Function test 2",
			args{
				matrix: [][]int{
					{1, 2, 3},
				},
			},
			[]int{1, 2, 3},
		},
		{
			"Function test 2",
			args{
				matrix: [][]int{
					{1},
					{2},
					{3},
				},
			},
			[]int{1, 2, 3},
		},
		{
			"Function test 3",
			args{
				matrix: [][]int{
					{1},
				},
			},
			[]int{1},
		},
		{
			"Empty test",
			args{
				matrix: [][]int{},
			},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Answer(tt.args.matrix); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Answer() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
