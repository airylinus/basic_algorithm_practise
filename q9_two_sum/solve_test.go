package q9_two_sum

import (
	"reflect"
	"testing"
)

func Test_towSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6},
				target: 9,
			},
			want: []int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := towSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("towSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
