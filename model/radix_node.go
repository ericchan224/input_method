package model

// radixNode 结构体用于存储 Radix 树中的节点
type radixNode struct {
	// 当前节点的相对路径
	path string
	// 完整路径
	fullPath string
	// 每个 indice 字符对应一个孩子节点的 path 首字母
	indices string
	// 后继节点
	children []*radixNode
	// 表示所有后继节点的汉字数量
	childrenCount int
	// 是否有路径以当前节点为终点
	end bool
	// 只有当end为true时才会有数值，用于存储当前词典所有的汉字，已经做好排序
	values []Character
	// 记录有多少路径途径当前节点
	passCnt int
}

// insert 向 radix 树中插入一个节点
func (rn *radixNode) insert(word string, values []Character) {
	fullWord := word

	// 如果当前节点为 root，此之前没有注册过子节点，则直接插入并返回
	if rn.path == "" && len(rn.children) == 0 {
		rn.insertWord(word, word, values)
		return
	}

walk:
	for {
		// 获取到 word 和当前节点 path 的公共前缀长度
		i := commonPrefixLen(word, rn.path)
		// 只要公共前缀大于 0，则一定经过当前节点，需要累加 passCnt
		if i > 0 {
			rn.passCnt++
		}

		// 公共前缀小于当前节点的相对路径，需要对节点进行分解
		if i < len(rn.path) {
			// 需要进行节点切割
			child := radixNode{
				// 进行相对路径切分
				path: rn.path[i:],
				// 继承完整路径
				fullPath: rn.fullPath,
				// 当前节点的后继节点进行委托
				children: rn.children,
				indices:  rn.indices,
				end:      rn.end,
				values:   rn.values,
				// 传承给孩子节点时，需要把之前累加上的 passCnt 计数扣除
				passCnt: rn.passCnt - 1,
			}

			// 续接上孩子节点
			rn.indices = string(rn.path[i])
			rn.children = []*radixNode{&child}
			// 调整原节点的 full path
			rn.fullPath = rn.fullPath[:len(rn.fullPath)-(len(rn.path)-i)]
			// 调整原节点的 path
			rn.path = rn.path[:i]
			// 原节点是新拆分出来的，目前不可能有单词以该节点结尾
			rn.end = false
			rn.values = nil
			rn.childrenCount += len(child.values)
		}

		// 公共前缀小于插入 word 的长度
		if i < len(word) {
			// 对 word 扣除公共前缀部分
			word = word[i:]
			// 获取 word 剩余部分的首字母
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				// 如果与后继节点还有公共前缀，则将 rn 指向子节点，然后递归执行流程
				if rn.indices[i] == c {
					rn = rn.children[i]
					continue walk
				}
			}

			// 到了这里，意味着 word 剩余部分与后继节点没有公共前缀了
			// 此时直接构造新的节点进行插入
			rn.indices += string(c)
			child := radixNode{}
			child.insertWord(word, fullWord, values)
			rn.children = append(rn.children, &child)
			rn.childrenCount += len(child.values)
			return
		}

		// 倘若公共前缀恰好是 path，需要将 end 置为 true
		rn.end = true
		rn.values = values
		return
	}
}

// 传入相对路径和完整路径，补充一个新生成的节点信息
func (rn *radixNode) insertWord(path, fullPath string, values []Character) {
	rn.path, rn.fullPath = path, fullPath
	rn.passCnt = 1
	rn.values = values
	rn.end = true
}

// search 函数在 radix 树中查找是否存在 word 对应的节点
func (rn *radixNode) search(word string) *radixNode {
walk:
	for {
		prefix := rn.path
		// word 长于 path
		if len(word) > len(prefix) {
			// 没匹配上，直接返回 nil
			if word[:len(prefix)] != prefix {
				return nil
			}
			// word 扣除公共前缀后的剩余部分
			word = word[len(prefix):]
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				// 后继节点还有公共前缀，继续匹配
				if c == rn.indices[i] {
					rn = rn.children[i]
					continue walk
				}
			}
			// word 还有剩余部分，但是 prefix 不存在后继节点和 word 剩余部分有公共前缀了
			// 必然不存在
			return nil
		}

		// 和当前节点精准匹配上了
		if word == prefix {
			return rn
		}

		// 走到这里意味着 len(word) <= len(prefix) && word != prefix
		return rn
	}
}