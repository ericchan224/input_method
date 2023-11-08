package model

import (
	"strings"
	// "icode.baidu.com/baidu/goodcoder/input_method/common"
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
		if len(node.values) < 10 {
			return node.values[:]
		}
		return node.values[:10]
	}

	// 表示不完全匹配
	// * 不同频次的汉字，频次越高的排在越前面
	// * 相同频次的汉字，根据对应的拼音的字母序排列，字母序越小的排在越前面。
	// * 相同频次的汉字，对应的拼音字母序也相同，则根据文件中的顺序排列。
	if node != nil && strings.HasPrefix(node.fullPath, word) {
		characters := make([]Character, 0, 10)
		stack := []*radixNode{node}
		for len(stack) > 0 {
			newStack := []*radixNode{}
			for _, n := range stack {
				if n.end {
					characters = append(characters, n.values...)
				}
				newStack = append(newStack, n.children...)
			}
			stack = newStack
		}
		// 按照规则排序
		// CharacterSort(characters)
		Sort(characters, 0, len(characters)-1)
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

// Erase 删除调一个字典
func (mim *MyInputMethod) Erase(word string) bool {
	if !mim.Search(word) {
		return false
	}

	// Root 直接精准命中了
	if mim.Root.fullPath == word {
		// 如果一个孩子都没有
		if len(mim.Root.indices) == 0 {
			mim.Root.path = ""
			mim.Root.fullPath = ""
			mim.Root.end = false
			mim.Root.passCnt = 0
			return true
		}

		// 如果只有一个孩子
		if len(mim.Root.indices) == 1 {
			mim.Root.children[0].path = mim.Root.path + mim.Root.children[0].path
			mim.Root = mim.Root.children[0]
			return true
		}

		// 如果有多个孩子
		for i := 0; i < len(mim.Root.indices); i++ {
			mim.Root.children[i].path = mim.Root.path + mim.Root.children[0].path
		}

		newRoot := radixNode{
			indices:  mim.Root.indices,
			children: mim.Root.children,
			passCnt:  mim.Root.passCnt - 1,
		}
		mim.Root = &newRoot
		return true
	}

	// 确定 word 存在的情况下
	move := mim.Root
	// Root 单独作为一个分支处理
	// 其他情况下，需要对孩子进行处理
walk:
	for {
		move.passCnt--
		prefix := move.path
		word = word[len(prefix):]
		c := word[0]
		for i := 0; i < len(move.indices); i++ {
			if move.indices[i] != c {
				continue
			}

			// 精准命中但是他仍有后继节点
			if move.children[i].path == word && move.children[i].passCnt > 1 {
				move.children[i].end = false
				move.children[i].passCnt--
				return true
			}

			// 找到对应的 child 了
			// 如果说后继节点的 passCnt = 1，直接干掉
			if move.children[i].passCnt > 1 {
				move = move.children[i]
				continue walk
			}

			move.children = append(move.children[:i], move.children[i+1:]...)
			move.indices = move.indices[:i] + move.indices[i+1:]
			// 如果干掉一个孩子后，发现只有一个孩子了，并且自身 end 为 false 则需要进行合并
			if !move.end && len(move.indices) == 1 {
				// 合并自己与唯一的孩子
				move.path += move.children[0].path
				move.fullPath = move.children[0].fullPath
				move.end = move.children[0].end
				move.indices = move.children[0].indices
				move.children = move.children[0].children
			}

			return true
		}
	}
}
