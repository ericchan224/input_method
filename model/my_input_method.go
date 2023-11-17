package model

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"icode.baidu.com/baidu/goodcoder/input_method/common"
)

// MyInputMethod ...
type MyInputMethod struct {
	// 你的数据结构在这里
	Root *radixNode
	// 锁
	sync.Mutex
}

// NewMyInputMethod 根据传入的词典文件创建一个新的输入法实例，
// 如果词典文件格式有误，忽略格式有误的文件
func NewMyInputMethod(dicts []string) *MyInputMethod {
	mim := MyInputMethod{
		Root: &radixNode{},
	}

	wg := &sync.WaitGroup{}
	for i := range dicts {
		dict := dicts[i]
		wg.Add(1)
		go func(dict string, wg *sync.WaitGroup) {
			defer wg.Done()

			dictArr := strings.Split(dict, "/")
			word := strings.Split(dictArr[len(dictArr)-1], ".")[0]
			wordLength := len(word)
			characters := []Character{}

			// 通过dict格式判断切换处理逻辑
			var res io.Reader
			switch dict[:4] {
			case "http":
				// 发起网络请求获取字典文件
				resp, err := common.HTTPCli.Get(dict)
				if err != nil {
					log.Println("http error:", err)
					return
				}
				defer resp.Body.Close()
				res = resp.Body
			default:
				// 直接从本地获取字典文件
				f, err := os.Open(dict)
				defer f.Close()
				if err != nil {
					log.Println("open file error:", err)
					return
				}
				res = f
			}

			// 解析返回内容
			reader := bufio.NewReader(res)
			var l int
			for {
				l++
				line, err := reader.ReadBytes('\n')
				if err != nil && err != io.EOF {
					log.Println("err:", err)
					continue
				}
				a := strings.TrimRight(string(line), "\n")
				if a != "" {
					c := strings.Split(a, " ")
					count, err := strconv.Atoi(c[1])
					if err != nil {
						log.Println("", err)
						continue
					}
					characters = append(characters, Character{
						Word:  c[0],
						Spell: word,
						Len:   wordLength,
						Count: count,
						Line:  l,
					})
				}
				if err == io.EOF {
					break
				}
			}

			// 插入Radix树
			// 按照频次排序
			CharacterSort(characters)
			mim.Lock()
			defer mim.Unlock()
			mim.Root.insert(word, characters)
		}(dict, wg)
	}
	wg.Wait()
	return &mim
}

// FindWords 根据输入的拼音返回对应的汉字，返回规则见功能描述
func (mim *MyInputMethod) FindWords(spell string) (words []string) {
	if spell == "" {
		return
	}
	// Your code here
	characters := mim.GetCharacter(spell)
	for _, v := range characters {
		words = append(words, v.Word)
	}
	return
}
