package q7_add_one

import (
	"reflect"
	"testing"
)

func Test_addOne(t *testing.T) {
	type args struct {
		nums []int
		want []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "[1,2,3]==>[2,3,4]",
			args: args{
				nums: []int{1, 2, 3},
				want: []int{2, 3, 4},
			},
		},
		{
			name: "[8,9,9]==>[9,1,0]",
			args: args{
				nums: []int{8, 9, 9},
				want: []int{9, 1, 0},
			},
		},
		{
			name: "[9,9,9]==>[9,1,0]",
			args: args{
				nums: []int{8, 9, 9},
				want: []int{9, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := addOne(tt.args.nums)
			for i := 0; i < len(tt.args.nums)-1; i++ {
				if r[i] != tt.args.nums[i] {
					t.Failed()
				}
			}
		})
	}
}

func Test_plusOne(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		wantR []int
	}{
		// TODO: Add test cases.
		{
			name:  "[1,2,3]==>[1,2,4]",
			args:  args{nums: []int{1, 2, 3}},
			wantR: []int{1, 2, 4},
		},
		{
			name:  "[9,9,9]==>[1,0,0,0]",
			args:  args{nums: []int{9, 9, 9}},
			wantR: []int{1, 0, 0, 0},
		},
		{
			name:  "[9,9,8]==>[9,9,9]",
			args:  args{nums: []int{9, 9, 8}},
			wantR: []int{9, 9, 9},
		},
		{
			name:  "[8,9,9,9]==>[9,0,0,0]",
			args:  args{nums: []int{8, 9, 9, 9}},
			wantR: []int{9, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := plusOne(tt.args.nums); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("plusOne() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
