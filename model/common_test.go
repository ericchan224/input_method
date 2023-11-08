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

// TestSort 用于测试排序算法
func TestSort(t *testing.T) {
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
	t.Run("quick sort", func(t *testing.T) {
		Sort(ch, 0, len(ch)-1)
		t.Logf("%v", ch)
	})
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

// TestCompare 是一个用于测试 Compare 函数的函数
func TestCompare(t *testing.T) {
	type args struct {
		a Character
		b Character
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "汉字",
			args: args{
				a: Character{
					Word:  "a",
					Len:   1,
					Count: 1,
				},
				b: Character{
					Word:  "a",
					Len:   1,
					Count: 1,
				},
			},
			want: true,
		},
		{
			name: "繁体字",
			args: args{
				a: Character{
					Word:  "A",
					Len:   1,
					Count: 1,
				},
				b: Character{
					Word:  "a",
					Len:   2,
					Count: 1,
				},
			},
			want: true,
		},
		{
			name: "繁字",
			args: args{
				a: Character{
					Word:  "a",
					Len:   2,
					Count: 2,
				},
				b: Character{
					Word:  "b",
					Len:   3,
					Count: 2,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%t", compare(tt.args.a, tt.args.b))
		})
	}
}

// TestPartition 函数对partition函数进行测试
func TestPartition(t *testing.T) {
	tests := []struct {
		in   []Character
		want int
	}{
		{
			in:   []Character{{Word: "a"}, {Word: "b"}, {Word: "c"}, {Word: "d"}, {Word: "e"}},
			want: 2,
		},
		{
			in:   []Character{{Word: "a"}, {Word: "b"}, {Word: "c"}, {Word: "d"}, {Word: "e"}, {Word: "f"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := partition(tt.in, 0, len(tt.in)-1)
		if got != tt.want {
			t.Fatalf("in: %v(MISSING), got: %v(MISSING), want: %v(MISSING)", tt.in, got, tt.want)
		}
	}
}
