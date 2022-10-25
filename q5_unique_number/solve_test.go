package q5_unique_number

import "testing"

func Test_getUniqueNumberVersion1(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		// TODO: Add test cases.
		{
			name: "test-001",
			args: args{
				numbers: []int{1, 2, 2, 33, 33, 44, 44},
			},
			wantNum: 1,
		},
		{
			name:    "test-002",
			args:    args{numbers: []int{44, 44, 13, 45, 13, 76, 76}},
			wantNum: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNum := getUniqueNumberVersion1(tt.args.numbers); gotNum != tt.wantNum {
				t.Errorf("getUniqueNumberVersion1() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}

func Test_getUniqueNumberVersion02(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test-002",
			args: args{numbers: []int{44, 44, 13, 45, 13, 76, 76}},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqueNumberVersion02(tt.args.numbers); got != tt.want {
				t.Errorf("getUniqueNumberVersion02() = %v, want %v", got, tt.want)
			}
		})
	}
}
