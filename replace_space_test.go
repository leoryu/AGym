package main

import "testing"

func TestReplaceSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantNewStr string
	}{
		{
			"Test 1",
			args{
				str: "abc 1 3",
			},
			"abc%201%203",
		},
		{
			"Test 2",
			args{
				str: "",
			},
			"",
		},
		{
			"Test 3",
			args{
				str: " ",
			},
			"%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewStr := ReplaceSpace(tt.args.str); gotNewStr != tt.wantNewStr {
				t.Errorf("ReplaceSpace() = %v, want %v", gotNewStr, tt.wantNewStr)
			}
		})
	}
}

