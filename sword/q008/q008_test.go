package q008

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			"Function test 1",
			args{
				input: []int{3, 4, 5, 1, 2},
			},
			[]int{1},
		},
		{
			"Function test 2",
			args{
				input: []int{2, 3, 4, 5, 1},
			},
			[]int{1},
		},
		{
			"Empty",
			args{
				input: []int{},
			},
			[]int{0},
		},
		{
			"All equal",
			args{
				input: []int{1, 1, 1},
			},
			[]int{1},
		},
		{
			"Just increase",
			args{
				input: []int{1, 2, 3},
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Answer(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Answer() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

