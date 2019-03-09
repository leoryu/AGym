package main

import (
	"testing"

	"github.com/leoryu/AGym/linkedlist"
)

func TestRecersePrint(t *testing.T) {
	head := &linkedlist.Single{
		Data: "head",
		Next: &linkedlist.Single{
			Data: "node",
			Next: &linkedlist.Single{
				Data: "last",
				Next: nil,
			},
		},
	}

	type args struct {
		head *linkedlist.Single
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			"Test 1",
			args{
				head: head,
			},
			`last
node
head
`,
		},
		{
			"Test 2",
			args{
				head: nil,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := RecersePrint(tt.args.head); gotResult != tt.wantResult {
				t.Errorf("RecersePrint() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

