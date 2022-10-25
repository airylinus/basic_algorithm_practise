package q8_move_zero

import (
	"reflect"
	"testing"
)

func Test_moveZero(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "[1,0,2,3,0,5,0,0,8]",
			args: args{nums: []int{1, 0, 2, 3, 0, 5, 0, 0, 8}},
			want: []int{1, 2, 3, 5, 8, 0, 0, 0, 0},
		},
		{
			name: "[1,0,2,3]->[1,2,3,0]",
			args: args{nums: []int{1, 0, 2, 3}},
			want: []int{1, 2, 3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZero(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
