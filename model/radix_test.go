package model

import (
	"sync"
	"testing"
)

var (
	radix *Radix
	onceT = &sync.Once{}
)

// GetRadix 返回一个单例的radix树
func GetRadix() *Radix {
	onceT.Do(func() {
		radix = NewRadix()
		radix.Insert("shui", []Character{
			{Word: "税", Len: 4, Count: 5},
			{Word: "谁", Len: 4, Count: 9},
			{Word: "水", Len: 4, Count: 10},
			{Word: "水", Len: 4, Count: 10},
		})
		radix.Insert("shu", []Character{
			{Word: "树", Len: 3, Count: 5},
			{Word: "数", Len: 3, Count: 9},
			{Word: "书", Len: 3, Count: 10},
		})
	})
	return radix
}

// TestInsert 测试Insert函数
func TestInsert(t *testing.T) {
	r := NewRadix()
	words := []string{"shui", "shu", "ch"}
	r.Insert(words[0], []Character{
		{Word: "税", Len: 4, Count: 5},
		{Word: "谁", Len: 4, Count: 9},
		{Word: "水", Len: 4, Count: 10},
	})

	for _, word := range words {
		t.Run(word, func(t *testing.T) {
			if ok := r.Search(word); ok {
				t.Logf("word %s is already insert: %t", word, ok)
			}
		})
	}
}

// TestRadixSearch 是一个用于测试Radix的函数
func TestRadixSearch(t *testing.T) {
	r := NewRadix()

	// test addWord
	r.Insert("chen", []Character{{Word: "陈", Len: 4, Count: 1}})

	words := []string{"chen", "ch"}
	for _, word := range words {
		t.Run(word, func(t *testing.T) {
			t.Logf("search word %s : %t", word, r.Search(word))
		})
	}
}

// TestGetCharacter 是一个测试函数，用于测试 GetCharacter 函数的正确性
func TestGetCharacter(t *testing.T) {
	rad := NewRadix()
	rad.Insert("chen", []Character{
		{Word: "陈", Len: 4, Count: 10},
		{Word: "晨", Len: 4, Count: 9},
		{Word: "趁", Len: 4, Count: 8},
	})
	rad.Insert("che", []Character{
		{Word: "车", Len: 3, Count: 10},
		{Word: "撤", Len: 3, Count: 7},
	})
	rad.Insert("cheng", []Character{
		{Word: "成", Len: 5, Count: 10},
		{Word: "城", Len: 5, Count: 7},
		{Word: "称", Len: 5, Count: 9},
	})

	spells := []string{"ch", "chen", "che", "cheng"}
	for _, spell := range spells {
		t.Run(spell, func(t *testing.T) {
			t.Logf("%s character %v:", spell, rad.GetCharacter(spell))
		})
	}
}

// TestStartWith 测试 Radix 树是否支持按前缀查询
func TestStartWith(t *testing.T) {
	rad := NewRadix()
	rad.Insert("chen", []Character{{Word: "陈", Len: 4, Count: 1}})

	spells := []string{"ch", "chen", "ji"}
	for _, spell := range spells {
		t.Run(spell, func(t *testing.T) {
			t.Logf("start with %s : %t:", spell, rad.StartWith(spell))
		})
	}
}

// TestRadix_Erase 测试 Radix 类型中 Erase 方法的实现是否正确
func TestRadix_Erase(t *testing.T) {
	r := NewRadix()
	r.Insert("chen", []Character{{Word: "陈", Len: 4, Count: 1}})

	t.Logf("search chen %t:", r.Search("chen"))

	found := r.Erase("chen")
	t.Logf("erase chen %t:", found)

	t.Logf("search chen %t:", r.Search("chen"))
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

// TestRadixRange1 是一个用于测试Radix树范围查找的函数
func TestRadixRange1(t *testing.T) {
	radix := NewRadix()
	a := Character{"a", 1, 1}
	radix.Insert(a.Word, []Character{a})
	b := Character{"B", 2, 1}
	radix.Insert(b.Word, []Character{b})

	t.Run("range", func(t *testing.T) {
		radix.Range()
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

// TestPassCnt 是一个用于测试 PassCnt 函数的 Go 语言测试函数
func TestPassCnt(t *testing.T) {
	r := GetRadix()
	t.Run("PassCnt", func(t *testing.T) {
		t.Logf("pass cnt: %d", r.PassCnt("shu"))
	})
}
