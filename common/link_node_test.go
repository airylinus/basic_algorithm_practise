package common

import (
	"reflect"
	"testing"
)

func TestMakeLinkNode(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want *LinkNode
	}{
		// TODO: Add test cases.
		{
			name: "test-001",
			args: args{in: []int{1, 2, 3, 4, 5}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeLinkNode(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Log(got.Print())
				t.Errorf("MakeLinkNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
