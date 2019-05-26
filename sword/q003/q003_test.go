package q003

import "testing"

func TestAnswer(t *testing.T) {
	type args struct {
		inputArray     [][]int
		expectedNumber int
	}
	tests := []struct {
		name          string
		args          args
		wantIsExisted bool
	}{
		{
			"Function test 1",
			args{
				inputArray: [][]int{
					[]int{1, 2, 8, 9},
					[]int{2, 4, 9, 12},
					[]int{4, 7, 10, 13},
					[]int{6, 8, 11, 15},
				},
				expectedNumber: 5,
			},
			false,
		},
		{
			"Function test 2",
			args{
				inputArray: [][]int{
					[]int{1, 2, 8, 9},
					[]int{2, 4, 9, 12},
					[]int{4, 7, 10, 13},
					[]int{6, 8, 11, 15},
				},
				expectedNumber: 7,
			},
			true,
		},
		{
			"Empty 1",
			args{
				inputArray:     [][]int{},
				expectedNumber: 0,
			},
			false,
		},
		{
			"Empty 2",
			args{
				inputArray:     nil,
				expectedNumber: 0,
			},
			false,
		},
		{
			"Just one number 1",
			args{
				inputArray: [][]int{
					[]int{1},
				},
				expectedNumber: 0,
			},
			false,
		},
		{
			"Just one number 2",
			args{
				inputArray: [][]int{
					[]int{1},
				},
				expectedNumber: 1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsExisted := Answer(tt.args.inputArray, tt.args.expectedNumber); gotIsExisted != tt.wantIsExisted {
				t.Errorf("Answer() = %v, want %v", gotIsExisted, tt.wantIsExisted)
			}
		})
	}
}
