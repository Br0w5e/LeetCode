package main

// 5519. 重新排列单词间的空格
import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {
	//记录空格个数
	space := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			space++
		}
	}
	//去掉空格
	parts := strings.Split(text, " ")
	rparts := make([]string, 0)
	for i := 0; i < len(parts); i++ {
		if parts[i] != "" {
			rparts = append(rparts, parts[i])
		}
	}

	size := len(rparts)
	//特殊情况：一个单词
	if size == 1 {
		return rparts[0] + strings.Repeat(" ", space)
	}
	//插入和末尾添加
	text_insert := strings.Repeat(" ", space/(size-1))
	text_add := strings.Repeat(" ", space%(size-1))
	//返回结果
	return strings.Join(rparts, text_insert) + text_add
}
func main() {
	text := "  this   is  a sentence "
	fmt.Println(reorderSpaces(text))
}
