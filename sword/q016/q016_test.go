package q016

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type args struct {
		head *NodeList
	}
	tests := []struct {
		name       string
		args       args
		wantTarget *NodeList
	}{
		{
			"Function test 1",
			args{
				head: &NodeList{
					Value: 1,
					NextNode: &NodeList{
						Value: 2,
						NextNode: &NodeList{
							Value:    3,
							NextNode: nil,
						},
					},
				},
			},
			&NodeList{
				Value: 3,
				NextNode: &NodeList{
					Value: 2,
					NextNode: &NodeList{
						Value:    1,
						NextNode: nil,
					},
				},
			},
		},
		{
			"Zero test",
			args{
				head: nil,
			},
			nil,
		},
		{
			"Test for single node",
			args{
				head: &NodeList{},
			},
			&NodeList{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTarget := Answer(tt.args.head); !reflect.DeepEqual(gotTarget, tt.wantTarget) {
				t.Errorf("Answer() = %v, want %v", gotTarget, tt.wantTarget)
			}
		})
	}
}
