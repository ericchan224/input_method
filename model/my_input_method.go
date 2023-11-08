package model

import (
	"bufio"
	"fmt"
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
	radix *Radix
	// 锁
	sync.RWMutex
}

// NewMyInputMethod 根据传入的词典文件创建一个新的输入法实例，
// 如果词典文件格式有误，忽略格式有误的文件
func NewMyInputMethod(dicts []string) *MyInputMethod {
	mim := MyInputMethod{}
	mim.SetRadix(NewRadix())

	wg := &sync.WaitGroup{}
	for i := range dicts {
		dict := dicts[i]
		go func(dict string, mim *MyInputMethod, wg *sync.WaitGroup) {
			wg.Add(1)
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
				defer resp.Body.Close()
				if err != nil {
					log.Println("http error:", err)
					return
				}
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
			for {
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
						Len:   wordLength,
						Count: count,
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
			mim.Insert(word, characters)
		}(dict, &mim, wg)
	}
	wg.Wait()
	return &mim
}

// Insert 向MyInputMethod中插入一个单词和对应的字符
func (mim *MyInputMethod) Insert(word string, characters []Character) {
	if mim.radix == nil {
		fmt.Printf("radix is nil")
		return
	}
	mim.radix.Insert(word, characters)
}

// SetRadix 将传入的 Radix 指针赋值给 MyInputMethod 结构体中的 radix 成员变量
func (mim *MyInputMethod) SetRadix(radix *Radix) {
	mim.radix = radix
}

// FindWords 根据输入的拼音返回对应的汉字，返回规则见功能描述
func (mim *MyInputMethod) FindWords(spell string) (words []string) {
	// Your code here
	characters := mim.radix.GetCharacter(spell)
	for _, v := range characters {
		words = append(words, v.Word)
	}
	return
}
