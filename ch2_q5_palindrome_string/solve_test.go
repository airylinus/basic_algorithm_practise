package ch2_q5_palindrome_string

import "testing"

func Test_palindromeString(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test-001_abcba",
			args: args{target: "abcba"},
			want: true,
		},
		{
			name: "test-001_abcba",
			args: args{target: "0123210"},
			want: true,
		},
		{
			name: "test-001_abcba",
			args: args{target: "xyz123321xxx"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromeString(tt.args.target); got != tt.want {
				t.Errorf("palindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
