package main

//1451. 重新排列句子中的单词
import (
	"sort"
	"strings"
)

//先转化为小写，使用SliceStable排序
func arrangeWords(text string) string {
	text = strings.ToLower(text)
	arr := strings.Split(text, " ")
	sort.SliceStable(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	res := []byte(strings.Join(arr, " "))
	res[0] = byte(res[0]-32)
	return string(res)
}
