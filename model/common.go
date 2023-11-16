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
			return characters[i].Len < characters[j].Len
		}

		return characters[i].Count > characters[j].Count
	})
}

// Sort 是一个递归函数，用于对字符数组ch进行排序，
func Sort(ch []Character, left, right int) {
	if left >= right {
		return
	}

	loc := partition(ch, left, right)
	Sort(ch, left, loc-1)
	Sort(ch, loc+1, right)
}

// partition 函数用于将字符数组ch[left, right]划分为两部分，使得左边的元素小于等于pivot，右边的元素大于pivot。
func partition(ch []Character, left, right int) int {
	pivot := ch[left]
	i := left + 1
	j := right

	for {
		for i <= right && compare(ch[i], pivot) {
			i++
		}
		for j > left && compare(pivot, ch[j]) {
			j--
		}

		if i >= j {
			break
		}

		ch[i], ch[j] = ch[j], ch[i]
		i++
		j--
	}

	ch[left], ch[j] = ch[j], ch[left]

	return j
}

// compare 函数接收两个 Character 类型的参数 a 和 b，返回一个布尔值。
func compare(a, b Character) bool {
	if a.Count == b.Count {
		return a.Len < b.Len
	}

	return a.Count > b.Count
}
