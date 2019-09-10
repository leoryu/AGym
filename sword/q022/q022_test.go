package q022

import "testing"

func TestAnswer(t *testing.T) {
	type args struct {
		in  []int
		out []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Function test 1",
			args{
				in:  []int{1, 2, 3, 4, 5},
				out: []int{5, 4, 3, 2, 1},
			},
			true,
		},
		{
			"Function test 2",
			args{
				in:  []int{1, 2, 3, 4, 5},
				out: []int{4, 5, 3, 2, 1},
			},
			true,
		},
		{
			"Function test 3",
			args{
				in:  []int{1, 2, 3, 4, 5},
				out: []int{4, 3, 5, 1, 2},
			},
			false,
		},
		{
			"Function test 4",
			args{
				in:  []int{1, 2, 3, 4, 5},
				out: []int{4, 3, 5},
			},
			false,
		},
		{
			"Function test 5",
			args{
				in:  []int{1, 2, 3},
				out: []int{4, 3, 5, 2, 1},
			},
			false,
		},
		{
			"Empty test 1",
			args{
				in:  []int{},
				out: []int{},
			},
			true,
		},
		{
			"Empty test 2",
			args{
				in:  []int{1},
				out: []int{},
			},
			false,
		},
		{
			"Empty test 3",
			args{
				in:  []int{},
				out: []int{1},
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Answer(tt.args.in, tt.args.out); got != tt.want {
				t.Errorf("Answer() = %v, want %v", got, tt.want)
			}
		})
	}
}
