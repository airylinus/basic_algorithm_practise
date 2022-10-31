package ch2_q9_longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		sources []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test-001_first/fish/fib",
			args: args{sources: []string{"first", "fish", "fib"}},
			want: "fi",
		},
		{
			name: "test-001_abcfirst/abcfirsh/abcdfirb",
			args: args{sources: []string{"abcfirst", "abcfirsh", "abcdfirb"}},
			want: "abc",
		},
		{
			name: "test-001_abcfirst/abcfirsh/abcdfirb",
			args: args{sources: []string{"abcfirst", "abcfirsh", "dfirb"}},
			want: "",
		},
		{
			name: "test-001_ab/a",
			args: args{sources: []string{"ab", "a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.sources); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
