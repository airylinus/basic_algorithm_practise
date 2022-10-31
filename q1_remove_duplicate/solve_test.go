package q1_remove_duplicate

import "testing"

func Test_removeDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test001-[1,2,3,3]",
			args: args{nums: []int{1, 2, 3, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
