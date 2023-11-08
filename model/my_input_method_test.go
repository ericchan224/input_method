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

	tests := []struct {
		name  string
		input *MyInputMethod
		spell string
		want  []string
	}{
		{
			name:  "test find words 4",
			input: input,
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

// TestMyInputMethod_Insert 是一个用于测试 MyInputMethod 结构体的 Insert 方法的函数
func TestMyInputMethod_Insert(t *testing.T) {
	t.Run("insert word:", func(t *testing.T) {
		mim := MyInputMethod{}
		mim.SetRadix(NewRadix())
		mim.Insert("zhan", []Character{
			{
				"站",
				1,
				10,
			},
		})
	})
}

// TestSetRadix 是一个用于测试SetRadix函数的测试函数
func TestSetRadix(t *testing.T) {
	t.Run("set radix:", func(t *testing.T) {
		mim := MyInputMethod{}
		mim.SetRadix(NewRadix())
	})
}
