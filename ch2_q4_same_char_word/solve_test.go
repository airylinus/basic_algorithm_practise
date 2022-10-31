package ch2_q4_same_char_word

import "testing"

func Test_isSameCharWord(t *testing.T) {
	type args struct {
		source string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test001-abcc==cbca",
			args: args{
				source: "abcc",
				target: "cbca",
			},
			want: true,
		},
		{
			name: "test001-abcc==cbca",
			args: args{
				source: "abccc",
				target: "cbcac",
			},
			want: true,
		},
		{
			name: "test001-abcc==cbca",
			args: args{
				source: "abcccx",
				target: "cbcaca",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameCharWord(tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("isSameCharWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
