package model

import (
	"testing"
)

func TestCommonPrefixLen(t *testing.T) {
	type args struct {
		wordA string
		wordB string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_01", args{"abc", "abcd"}, 3},
		{"Test_02", args{"abc", "ab"}, 2},
		{"Test_03", args{"", ""}, 0},
		{"Test_04", args{"abc", "abcde"}, 3},
		{"Test_05", args{"abcd", "abc"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonPrefixLen(tt.args.wordA, tt.args.wordB); got != tt.want {
				t.Errorf("commonPrefixLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
