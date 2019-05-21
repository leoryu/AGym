package main

import (
	"math/big"
	"testing"
)

func TestFibonacciSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name        string
		args        args
		wantNumberN *big.Int
	}{
		{
			"0",
			args{
				n: 0,
			},
			big.NewInt(0),
		},
		{
			"1",
			args{
				n: 1,
			},
			big.NewInt(1),
		},
		{
			"10",
			args{
				n: 10,
			},
			big.NewInt(55),
		},
		{
			"17",
			args{
				n: 17,
			},
			big.NewInt(1597),
		},
		{
			"24",
			args{
				n: 24,
			},
			big.NewInt(46368),
		},
		{
			"50",
			args{
				n: 50,
			},
			big.NewInt(12586269025),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumberN := FibonacciSequence(tt.args.n); gotNumberN.Cmp(tt.wantNumberN) != 0 {
				t.Errorf("FibonacciSequence() = %v, want %v", gotNumberN, tt.wantNumberN)
			}
		})
	}
}

