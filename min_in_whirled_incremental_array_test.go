package main

import "testing"

func TestMinInWhirledIncrementalArray(t *testing.T) {
	type args struct {
		incrementalArray []int
	}
	tests := []struct {
		name          string
		args          args
		wantMinNumber int
	}{
		{
			"array 34512",
			args{
				incrementalArray: []int{3, 4, 5, 1, 2},
			},
			1,
		},
		{
			"array 12345",
			args{
				incrementalArray: []int{1, 2, 3, 4, 5},
			},
			1,
		},
		{
			"array 1",
			args{
				incrementalArray: []int{1},
			},
			1,
		},
		{
			"array 2341",
			args{
				incrementalArray: []int{2, 3, 4, 1},
			},
			1,
		},
		{
			"array 111",
			args{
				incrementalArray: []int{1, 1, 1, 1},
			},
			1,
		},
		{
			"array 345123",
			args{
				incrementalArray: []int{3, 4, 5, 1, 2, 3},
			},
			1,
		},
		{
			"array 33313",
			args{
				incrementalArray: []int{3, 3, 3, 1, 3},
			},
			1,
		},
		{
			"array 31333",
			args{
				incrementalArray: []int{3, 1, 3, 3, 3},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMinNumber := MinInWhirledIncrementalArray(tt.args.incrementalArray); gotMinNumber != tt.wantMinNumber {
				t.Errorf("MinInWhirledIncrementalArray() = %v, want %v", gotMinNumber, tt.wantMinNumber)
			}
		})
	}
}

