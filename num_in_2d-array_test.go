package main

import "testing"

func TestIsNumIn2dArray(t *testing.T) {
	type args struct {
		array [][]int
		num   int
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
	}{
		{
			"Test 1",
			args{
				array: [][]int{
					{1, 2, 3},
					{2, 3, 4},
				},
				num: 4,
			},
			true,
		},
		{
			"Test 2",
			args{
				array: [][]int{
					{1, 2, 3},
					{2, 3, 4},
				},
				num: 0,
			},
			false,
		},
		{
			"Test 3",
			args{
				array: [][]int{},
				num:   0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExist := IsNumIn2dArray(tt.args.array, tt.args.num); gotExist != tt.wantExist {
				t.Errorf("IsNumIn2dArray() = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
