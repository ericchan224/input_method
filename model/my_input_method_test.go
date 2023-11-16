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

	t.Run("test new my input method", func(t *testing.T) {
		_ = NewMyInputMethod(dicts)

	})
}

// TestMyInputMethod_FindWords 函数测试MyInputMethod的FindWords方法
func TestFindWords(t *testing.T) {
	dicts := []string{"../dict/zhan.dat",
		"../dict/zhang.dat",
		"../dict/de.dat",
		"../dict/chang.dat",
		"../dict/cha.dat",
		"../dict/chao.dat"}
	input := NewMyInputMethod(dicts)
	t.Logf("words: %v", input.FindWords("de"))
}
