package q012

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAnswer(t *testing.T) {
	type args struct {
		n int
	}
	wantForFunctionTest2 := []string{98: "99"}
	for i := 0; i < 99; i++ {
		wantForFunctionTest2[i] = strconv.Itoa(i + 1)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Function test 1",
			args{
				1,
			},
			[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			"Function test 2",
			args{
				2,
			},
			wantForFunctionTest2,
		},
		{
			"0 test",
			args{
				0,
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Answer(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Answer() = %v, want %v", got, tt.want)
			}
		})
	}
}

