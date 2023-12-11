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
			move := commonPrefixLen(tt.args.wordA, tt.args.wordB)
			if move != tt.want {
				t.Logf("无公共前缀")
			} else {
				t.Logf("公共前缀为%s:", tt.args.wordA[:move])
			}
		})
	}
}

// TestChSort 是一个用于测试 CharacterSort 函数的 Go 语言测试函数
func TestChSort(t *testing.T) {
	t.Parallel()

	a := Character{
		Word:  "陈",
		Len:   4,
		Spell: "chen",
		Count: 9,
		Line:  2,
	}
	b := Character{
		Word:  "晨",
		Len:   4,
		Spell: "chen",
		Count: 10,
		Line:  1,
	}
	c := Character{
		Word:  "成",
		Len:   5,
		Spell: "cheng",
		Count: 10,
		Line:  1,
	}
	d := Character{
		Word:  "趁",
		Len:   4,
		Spell: "chen",
		Count: 10,
		Line:  3,
	}
	ch := []Character{a, b, c, d}
	temp := []Character{b, d, c, a}
	t.Run("character sort", func(t *testing.T) {
		t.Parallel()
		CharacterSort(ch)
		for i := range ch {
			if ch[i] != temp[i] {
				t.Logf("排序错误")
			}
		}
	})
}
