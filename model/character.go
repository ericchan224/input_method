package model

// Character 结构体用于存储汉字的拼音和频次
type Character struct {
	// 汉字的值
	Word string
	// 拼音
	Spell string
	// 汉字拼音长度
	Len int
	// 汉字频次
	Count int
	// 文件中原来的顺序
	Line int
}
