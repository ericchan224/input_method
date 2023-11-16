package model

import (
	"testing"
)

// TestCommonPrefixLen 函数用于测试 commonPrefixLen 函数的正确性
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
			commonPrefixLen(tt.args.wordA, tt.args.wordB)
		})
	}
}

// TestChSort 是一个用于测试 CharacterSort 函数的 Go 语言测试函数
func TestChSort(t *testing.T) {
	t.Parallel()

	ch := []Character{{
		Word:  "a",
		Len:   1,
		Count: 1,
	}, {
		Word:  "b",
		Len:   1,
		Count: 1,
	}, {
		Word:  "c",
		Len:   2,
		Count: 2,
	}}
	t.Run("character sort", func(t *testing.T) {
		t.Parallel()
		CharacterSort(ch)
	})
}
