package model

import (
	"sync"
	"testing"
)

var (
	m     *MyInputMethod
	onceT = &sync.Once{}
)

// GetRadix 返回一个单例的radix树
func GetMethod() *MyInputMethod {
	onceT.Do(func() {
		m = &MyInputMethod{
			Root: &radixNode{},
		}
		m.Root.insert("shui", []Character{
			{Word: "税", Len: 4, Count: 5, Line: 4},
			{Word: "谁", Len: 4, Count: 9, Line: 3},
			{Word: "水", Len: 4, Count: 10, Line: 2},
			{Word: "水", Len: 4, Count: 10, Line: 1},
		})
		m.Root.insert("shu", []Character{
			{Word: "树", Len: 3, Count: 5, Line: 3},
			{Word: "数", Len: 3, Count: 9, Line: 2},
			{Word: "书", Len: 3, Count: 10, Line: 1},
		})
		m.Root.insert("she", []Character{
			{"舍", "she", 3, 10, 1},
		})
		m.Root.insert("shen", []Character{
			{"神", "shen", 4, 9, 1},
		})
		m.Root.insert("de", []Character{
			{"的", "de", 2, 10, 1},
		})
		m.Root.insert("chao", []Character{
			{"潮", "chao", 4, 10, 14},
			{"炒", "chao", 4, 10, 13},
			{"超", "chao", 4, 10, 12},
			{"吵", "chao", 4, 10, 11},
			{"钞", "chao", 4, 10, 10},
			{"巢", "chao", 4, 10, 9},
			{"朝", "chao", 4, 10, 8},
			{"晁", "chao", 4, 10, 7},
			{"嘲", "chao", 4, 10, 6},
			{"绰", "chao", 4, 10, 5},
			{"怊", "chao", 4, 10, 4},
			{"焯", "chao", 4, 10, 3},
			{"耖", "chao", 4, 10, 2},
			{"剿", "chao", 4, 10, 1},
		})
		m.Root.insert("zhang", []Character{
			{"长", "zhang", 5, 10, 14},
		})
		m.Root.insert("chang", []Character{
			{"长", "chang", 5, 10, 14},
		})
	})
	return m
}

// TestPassCnt 是一个用于测试 PassCnt 函数的 Go 语言测试函数
func TestPassCnt(t *testing.T) {
	r := GetMethod()

	spells := []struct {
		Spell string
		Cnt   int
	}{
		{"shui", 1},
		{"shu", 2},
		{"she", 2},
		{"shen", 1},
		{"de", 1},
		{"", 0},
		{"qiong", 0},
	}
	for _, v := range spells {
		t.Run("PassCnt", func(t *testing.T) {
			t.Logf("pass cnt: %v", v.Cnt == r.PassCnt(v.Spell))
		})
	}
}

// TestMyInputMethod_Search 是一个用于测试MyInputMethod的Search方法的函数
func TestMyInputMethod_Search(t *testing.T) {
	r := GetMethod()
	t.Run("test search", func(t *testing.T) {
		if !r.Search("shu") {
			t.Logf("word shu not founr, maybe shu is not a word")
		}
	})
}

// TestGetCharacter 是一个测试函数，用于获取字符。
func TestGetCharacter(t *testing.T) {
	m := GetMethod()
	spells := []string{"shui", "shu", "she", "shen", "de", "sh", "chao", "ch"}
	for _, v := range spells {
		t.Run("test get character", func(t *testing.T) {
			t.Logf("get character: %v", m.GetCharacter(v))
		})
	}
}

// TestStartWith 是一个用于测试StartWith函数的函数
func TestStartWith(t *testing.T) {
	m := GetMethod()
	t.Run("test start with", func(t *testing.T) {
		t.Logf("%v", m.StartWith("sh"))
	})
}
