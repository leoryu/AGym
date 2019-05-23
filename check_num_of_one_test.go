package main

import (
	"math"
	"testing"
)

func TestCheckNumOne(t *testing.T) {
	type args struct {
		checkedNum int
	}
	tests := []struct {
		name         string
		args         args
		wantNumOfOne int
	}{
		{
			"0",
			args{
				checkedNum: 0,
			},
			0,
		},
		{
			"1",
			args{
				checkedNum: 1,
			},
			1,
		},
		{
			"100",
			args{
				checkedNum: 100,
			},
			3,
		},
		{
			"-100",
			args{
				checkedNum: -100,
			},
			3,
		},
		{
			"MaxInt64",
			args{
				checkedNum: math.MaxInt64,
			},
			63,
		},
		{
			"MinInt64",
			args{
				checkedNum: math.MinInt64,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumOfOne := CheckNumOne(tt.args.checkedNum); gotNumOfOne != tt.wantNumOfOne {
				t.Errorf("CheckNumOne() = %v, want %v", gotNumOfOne, tt.wantNumOfOne)
			}
		})
	}
}

