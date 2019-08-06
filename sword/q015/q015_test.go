package q015

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type args struct {
		head *NodeList
		k    int
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
				k: 1,
			},
			&NodeList{
				Value:    3,
				NextNode: nil,
			},
		},
		{
			"Function test 2",
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
				k: 2,
			},
			&NodeList{
				Value: 2,
				NextNode: &NodeList{
					Value:    3,
					NextNode: nil,
				},
			},
		},
		{
			"Test for not long enough",
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
				k: 10,
			},
			nil,
		},
		{
			"Zero test",
			args{
				head: nil,
				k:    1,
			},
			nil,
		},
		{
			"Test for illegale k",
			args{
				head: &NodeList{},
				k:    -1,
			},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTarget := Answer(tt.args.head, tt.args.k); !reflect.DeepEqual(gotTarget, tt.wantTarget) {
				t.Errorf("Answer() = %v, want %v", gotTarget, tt.wantTarget)
			}
		})
	}
}

