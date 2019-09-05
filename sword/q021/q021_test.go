package q021

import (
	"reflect"
	"testing"
)

func intPointer(i int) *int {
	return &i
}

func TestStack_Top(t *testing.T) {
	fTest2 := Stack{}
	fTest2.Push(2)
	fTest2.Push(3)
	fTest3 := fTest2
	fTest3.Pop()
	emptyTest := fTest3
	emptyTest.Pop()
	emptyTest.Pop()
	type fields struct {
		stack []item
	}
	tests := []struct {
		name    string
		fields  fields
		wantTop *int
	}{
		{
			"Function test 1",
			fields{
				stack: []item{item{value: 1}, item{value: 2}},
			},
			intPointer(2),
		},
		{
			"Function test 2",
			fields{
				stack: fTest2.stack,
			},
			intPointer(3),
		},
		{
			"Function test 3",
			fields{
				stack: fTest3.stack,
			},
			intPointer(2),
		},
		{
			"Empty test",
			fields{
				stack: emptyTest.stack,
			},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			if gotTop := s.Top(); !reflect.DeepEqual(gotTop, tt.wantTop) {
				t.Errorf("Stack.Top() = %v, want %v", *gotTop, *tt.wantTop)
			}
		})
	}
}

func TestStack_Min(t *testing.T) {
	fTest1 := Stack{}
	fTest1.Push(0)
	fTest1.Push(2)
	fTest2 := fTest1
	fTest2.Push(-2)
	fTest3 := fTest2
	fTest3.Pop()
	fTest3.Pop()
	emptyTest := fTest3
	emptyTest.Pop()
	type fields struct {
		stack []item
	}
	tests := []struct {
		name    string
		fields  fields
		wantMin *int
	}{
		{
			"Function test 1",
			fields{stack: fTest1.stack},
			intPointer(0),
		},
		{
			"Function test 2",
			fields{stack: fTest2.stack},
			intPointer(-2),
		},
		{
			"Function test 3",
			fields{stack: fTest3.stack},
			intPointer(0),
		},
		{
			"Empty test 3",
			fields{stack: emptyTest.stack},
			nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			if gotMin := s.Min(); !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("Stack.Min() = %v, want %v", &gotMin, &tt.wantMin)
			}
		})
	}
}
