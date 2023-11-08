// Copyright 2023 Baidu Inc. All rights reserved.
// Use of this source code is governed by a xxx
// Use of this source code is governed by a xxx
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.

// Package main is special.  It defines a
// Package main is special.  It defines a
// standalone executable program, not a library.
// standalone executable program, not a library.
// Within package main the function main is also
// Within package main the function main is also
// special—it’s where execution of the program begins.
// special—it’s where execution of the program begins.
// Whatever main does is what the program does.
// Whatever main does is what the program does.
package main

import (
	// 第三方库
	// 项目自己
	// GO标准库
	"bufio"
	"fmt"
	"os"
	"strings"

	"icode.baidu.com/baidu/goodcoder/input_method/model"
)

// InputMethod 是一个输入方法接口，用于查找词组。
type InputMethod interface {
	FindWords(string) []string
}

// loop 是一个无限循环函数，从标准输入中读取词语并使用给定的 InputMethod 查找词语。
func loop(im InputMethod) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入词语：")
		spell, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		spell = strings.TrimRight(spell, "\n")
		if spell == "" {
			continue
		}
		words := im.FindWords(spell)
		fmt.Printf("结果展示：%s\n", strings.Join(words, ", "))
	}
}

// main 函数是程序的入口点
func main() {
	fmt.Println("Hello, World!")
	urls := os.Args[1:]
	im := model.NewMyInputMethod(urls)

	loop(im)
}
