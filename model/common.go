package model

import "sort"

// 求取两个单词的公共前缀
func commonPrefixLen(wordA, wordB string) int {
	var move int
	for move < len(wordA) && move < len(wordB) && wordA[move] == wordB[move] {
		move++
	}
	return move
}

// CharacterSort 废弃
func CharacterSort(characters []Character) {
	sort.Slice(characters, func(i, j int) bool {
		if characters[i].Count == characters[j].Count && characters[i].Len == characters[j].Len {
			return characters[i].Line < characters[j].Line
		}

		if characters[i].Count == characters[j].Count {
			return characters[i].Spell < characters[j].Spell
		}

		return characters[i].Count > characters[j].Count
	})
}
