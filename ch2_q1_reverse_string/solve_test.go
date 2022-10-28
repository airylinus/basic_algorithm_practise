package ch2_q1_reverse_string

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name  string
		args  args
		wantR []byte
	}{
		{
			name: "reverseString-001",
			args: args{
				s: []byte{'a', 'b', 'c'},
			},
			wantR: []byte{'c', 'b', 'a'},
		},
		{
			name: "reverseString-001",
			args: args{
				s: []byte{'a', 'b', 'c', 'd'},
			},
			wantR: []byte{'d', 'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := reverseString(tt.args.s); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("reverseString() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
