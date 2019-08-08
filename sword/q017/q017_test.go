package q017

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type args struct {
		l1 *NodeList
		l2 *NodeList
	}
	tests := []struct {
		name   string
		args   args
		wantMl *NodeList
	}{
		{
			"Function test 1",
			args{
				l1: &NodeList{Value: 1, NextNode: nil},
				l2: &NodeList{Value: 2, NextNode: nil},
			},
			&NodeList{Value: 1,
				NextNode: &NodeList{
					Value: 2, NextNode: nil,
				}},
		},
		{
			"Function test 2",
			args{
				l1: &NodeList{Value: 1, NextNode: &NodeList{
					Value:    3,
					NextNode: nil,
				}},
				l2: &NodeList{Value: 2, NextNode: nil},
			},
			&NodeList{Value: 1,
				NextNode: &NodeList{
					Value: 2, NextNode: &NodeList{
						Value:    3,
						NextNode: nil,
					},
				}},
		},
		{
			"Function test 3",
			args{
				l1: nil,
				l2: &NodeList{Value: 2, NextNode: nil},
			},
			&NodeList{Value: 2,
				NextNode: nil,
			},
		},
		{
			"Zero test",
			args{
				l1: nil,
				l2: nil,
			},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMl := Answer(tt.args.l1, tt.args.l2); !reflect.DeepEqual(gotMl, tt.wantMl) {
				t.Errorf("Answer() = %v, want %v", gotMl, tt.wantMl)
			}
		})
	}
}

