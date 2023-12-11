package model

import (
	"strings"
)

// Search 查看一个单词在 radix 当中是否存在
func (mim *MyInputMethod) Search(word string) bool {
	node := mim.Root.search(word)
	return node != nil && node.fullPath == word && node.end
}

// GetCharacter 获取汉字，并根据规则排序
func (mim *MyInputMethod) GetCharacter(word string) []Character {
	node := mim.Root.search(word)

	// 表示完全匹配
	if node != nil && node.fullPath == word && node.end {
		return node.values[:]
	}

	// 表示不完全匹配
	// * 不同频次的汉字，频次越高的排在越前面
	// * 相同频次的汉字，根据对应的拼音的字母序排列，字母序越小的排在越前面。
	// * 相同频次的汉字，对应的拼音字母序也相同，则根据文件中的顺序排列。
	if node != nil && strings.HasPrefix(node.fullPath, word) {
		characters := make([]Character, 0, 10)
		cMap := make(map[string]Character)
		stack := []*radixNode{node}
		for len(stack) > 0 {
			newStack := []*radixNode{}
			for _, n := range stack {
				if n.end {
					// 多音字过滤，保留频次最高的那个
					for _, c := range n.values {
						if _, ok := cMap[c.Word]; !ok {
							cMap[c.Word] = c
							continue
						}
						if c.Count > cMap[c.Word].Count {
							cMap[c.Word] = c
						}
					}
				}
				newStack = append(newStack, n.children...)
			}
			stack = newStack
		}
		for _, c := range cMap {
			characters = append(characters, c)
		}
		// 按照规则排序
		CharacterSort(characters)
		if len(characters) < 10 {
			return characters[:]
		}
		return characters[:10]
	}
	return nil
}

// StartWith 前缀匹配流程
func (mim *MyInputMethod) StartWith(prefix string) bool {
	node := mim.Root.search(prefix)
	return node != nil && strings.HasPrefix(node.fullPath, prefix)
}

// PassCnt 返回以 prefix 为前缀的路由对应的 passCnt 值
func (mim *MyInputMethod) PassCnt(prefix string) int {
	node := mim.Root.search(prefix)
	if node == nil || !strings.HasPrefix(node.fullPath, prefix) {
		return 0
	}
	return node.passCnt
}
