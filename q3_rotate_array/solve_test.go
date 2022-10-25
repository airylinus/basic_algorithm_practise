package q3_rotate_array

import (
	"reflect"
	"testing"
)

func Test_revert_nums(t *testing.T) {
	type args struct {
		nums  []int
		wants []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test-001",
			args: args{
				nums:  []int{1, 2, 3, 4, 5},
				wants: []int{5, 4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := revertNums(tt.args.nums)
			if len(r) != len(tt.args.nums) {
				t.Fail()
			}
			for i, _ := range r {
				if r[i] != tt.args.nums[i] {
					t.Fail()
				}
			}
			t.Log(r)
		})
	}
}

func Test_rotateArraySolve1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		args  args
		wantR []int
	}{
		// TODO: Add test cases.
		{
			name: "1,2,3,4,5,6,7 --> 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			wantR: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := rotateArraySolve1(tt.args.nums, tt.args.k); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("rotateArraySolve1() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
