package model

import (
	"testing"
)

// TestNewMyInputMethod 是一个用于测试 NewMyInputMethod 函数的函数
func TestNewMyInputMethod(t *testing.T) {
	dicts := []string{"../dict/zhan.dat",
		"../dict/zhang.dat",
		"../dict/de.dat",
		"../dict/chang.dat",
		"../dict/cha.dat",
		"../dict/chao.dat"}

	// t.Run("test new my input method", func(t *testing.T) {
	_ = NewMyInputMethod(dicts)

	// r.Radix.Range()
	// })
}

// TestMyInputMethod_FindWords 函数测试MyInputMethod的FindWords方法
func TestFindWords(t *testing.T) {
	// dicts := []string{"../dict/zhan.dat",
	// 	"../dict/zhang.dat",
	// 	"../dict/de.dat",
	// 	"../dict/chang.dat",
	// 	"../dict/cha.dat",
	// 	"../dict/chao.dat"}
	// input := NewMyInputMethod(dicts)
	// for int(DictCount.Count) < len(dicts) {
	// 	// 设置超时时间
	// 	time.Sleep(time.Second)
	// 	t.Log("waiting for dicts loading...")
	// }
	tests := []struct {
		name  string
		input *MyInputMethod
		spell string
		want  []string
	}{
		{
			name:  "test find words 4",
			input: nil,
			spell: "de",
		},
	}
	for _, tt := range tests {
		if tt.input == nil {
			continue
		}
		t.Logf("words: %v", tt.input.FindWords(tt.spell))
	}
}
