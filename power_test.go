package main

import "testing"

func TestCustomPower(t *testing.T) {
	type args struct {
		base     float64
		exponent int
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{
			"2,3",
			args{
				base: 2,
				exponent: 3,
			},
			8,
		},
		{
			"0,3",
			args{
				base: 0,
				exponent: 3,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CustomPower(tt.args.base, tt.args.exponent); gotResult != tt.wantResult {
				t.Errorf("CustomPower() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

