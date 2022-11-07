package ch3_q1_remove_linknode

import (
	"testing"

	"github.com/airylinus/basic_algorithm_practise/common"
)

func Test_removeLinkNode(t *testing.T) {
	type args struct {
		head      *common.LinkNode
		targetVal int
	}
	tests := []struct {
		name        string
		args        args
		wantNewHead *common.LinkNode
		wantResult  string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				head:      common.MakeLinkNode([]int{1, 2, 3, 5, 6, 7}),
				targetVal: 3,
			},
			wantNewHead: nil,
			wantResult:  "1 -> 2 -> 5 -> 6 -> 7",
		},
		{
			name: "test",
			args: args{
				head:      common.MakeLinkNode([]int{1, 2, 3, 5, 6, 7}),
				targetVal: 6,
			},
			wantNewHead: nil,
			wantResult:  "1 -> 2 -> 3 -> 5 -> 7",
		},
		{
			name: "test",
			args: args{
				head:      common.MakeLinkNode([]int{1, 2, 3, 5, 6, 7}),
				targetVal: 2,
			},
			wantNewHead: nil,
			wantResult:  "1 -> 3 -> 5 -> 6 -> 7",
		},
		{
			name: "test-4",
			args: args{
				head:      common.MakeLinkNode([]int{2, 3, 5, 6, 7}),
				targetVal: 2,
			},
			wantNewHead: nil,
			wantResult:  "3 -> 5 -> 6 -> 7",
		},
		{
			name: "test-4",
			args: args{
				head:      common.MakeLinkNode([]int{2, 3, 5, 6, 7}),
				targetVal: 6,
			},
			wantNewHead: nil,
			wantResult:  "2 -> 3 -> 5 -> 7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewHead := removeLinkNode(tt.args.head, tt.args.targetVal)
			t.Log(gotNewHead.Print())
			if gotNewHead.Print() != tt.wantResult {
				t.Errorf("removeLinkNode() = %v, want %v", gotNewHead.Print(), tt.wantResult)
			}
		})
	}
}
