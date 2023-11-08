package model

import (
	"testing"
)

// TestRadixNode_Insert 测试 RadixNode 的 Insert 方法
func TestRadixNode_Insert(t *testing.T) {
	type args struct {
		word  string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "insert test",
			args: args{
				word:  "test",
				value: "value",
			},
		},
		{
			name: "insert test",
			args: args{
				word:  "tester",
				value: "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &radixNode{}
			rn.insert(tt.args.word, []Character{})
			if len(rn.children[0].values) != 1 {
				t.Errorf("insert failed")
			}
		})
	}
}

// TestRadixNode_insertWord 用于测试 RadixNode 结构体的 insertWord 方法
func TestRadixNode_insertWord(t *testing.T) {
	type args struct {
		path     string
		fullPath string
		values   []Character
	}
	tests := []struct {
		name string
		args args
		want []Character
	}{
		{
			name: "test1",
			args: args{
				path:     "a",
				fullPath: "a",
				values:   []Character{},
			},
			want: []Character{},
		},
		{
			name: "test2",
			args: args{
				path:     "ab",
				fullPath: "ab",
				values:   []Character{},
			},
			want: []Character{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &radixNode{}
			rn.insertWord(tt.args.path, tt.args.fullPath, tt.args.values)
		})
	}
}

// TestradixNode_search 是一个测试函数，用于测试 radixNode 结构体的 search 方法。
func TestRadixNode_search(t *testing.T) {
	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want *radixNode
	}{
		{
			name: "test",
			args: args{
				word: "test",
			},
			want: &radixNode{path: "/test/"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rn := &radixNode{}
			rn.search(tt.args.word)
		})
	}
}
