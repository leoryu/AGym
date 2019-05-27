package q004

import "testing"

func TestAnswer(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name          string
		args          args
		wantOutputStr string
	}{
		{
			"Function test 1",
			args{
				inputStr: "We Are Happy",
			},
			"We%20Are%20Happy",
		},
		{
			"Empty",
			args{
				inputStr: "",
			},
			"",
		},
		{
			"Just space",
			args{
				inputStr: " ",
			},
			"%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputStr := Answer(tt.args.inputStr); gotOutputStr != tt.wantOutputStr {
				t.Errorf("Answer() = %v, want %v", gotOutputStr, tt.wantOutputStr)
			}
		})
	}
}

