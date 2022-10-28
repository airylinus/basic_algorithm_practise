package ch2_q2_first_unique_char

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestFirstUniqChar-001",
			args: args{s: "aabbccd"},
			want: 6,
		},
		{
			name: "TestFirstUniqChar-001",
			args: args{s: "xzy"},
			want: 0,
		},
		{
			name: "TestFirstUniqChar-001",
			args: args{s: "xzyxyz"},
			want: -1,
		},
		{
			name: "TestFirstUniqChar-001",
			args: args{s: "xzzxzzbba"},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
