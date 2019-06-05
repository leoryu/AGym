package q011

import "testing"

func TestAnswer(t *testing.T) {
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
				base:     2,
				exponent: 3,
			},
			8,
		},
		{
			"0,3",
			args{
				base:     0,
				exponent: 3,
			},
			0,
		},
		{
			"1,0",
			args{
				base:     1,
				exponent: 0,
			},
			1,
		},
		{
			"3,-1",
			args{
				base:     3,
				exponent: -1,
			},
			1.0 / 3,
		},
		{
			"2,-3",
			args{
				base:     2,
				exponent: -3,
			},
			1.0 / 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Answer(tt.args.base, tt.args.exponent); gotResult != tt.wantResult {
				t.Errorf("Answer() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

