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
			{Word: "税", Len: 4, Count: 5},
			{Word: "谁", Len: 4, Count: 9},
			{Word: "水", Len: 4, Count: 10},
			{Word: "水", Len: 4, Count: 10},
		})
		m.Root.insert("shu", []Character{
			{Word: "树", Len: 3, Count: 5},
			{Word: "数", Len: 3, Count: 9},
			{Word: "书", Len: 3, Count: 10},
		})
		m.Root.insert("she", []Character{
			{"舍", 3, 10},
		})
		m.Root.insert("shen", []Character{
			{"神", 4, 9},
		})
		m.Root.insert("de", []Character{
			{"的", 2, 10},
		})
		m.Root.insert("", []Character{
			{"潮", 4, 10},
			{"炒", 4, 10},
			{"超", 4, 10},
			{"吵", 4, 10},
			{"钞", 4, 10},
			{"巢", 4, 10},
			{"朝", 4, 10},
			{"晁", 4, 10},
			{"嘲", 4, 10},
			{"绰", 4, 10},
			{"怊", 4, 10},
			{"焯", 4, 10},
			{"耖", 4, 10},
			{"剿", 4, 10},
		})
	})
	return m
}

// TestPassCnt 是一个用于测试 PassCnt 函数的 Go 语言测试函数
func TestPassCnt(t *testing.T) {
	r := GetMethod()

	spells := []string{"shui", "shu", "she", "shen", "de", "", "qiong"}
	for _, v := range spells {
		t.Run("PassCnt", func(t *testing.T) {
			t.Logf("pass cnt: %d", r.PassCnt(v))
		})
	}
}

// TestMyInputMethod_Search 是一个用于测试MyInputMethod的Search方法的函数
func TestMyInputMethod_Search(t *testing.T) {
	r := GetMethod()
	t.Run("test search", func(t *testing.T) {
		r.Search("shu")
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
