package main

import (
	"fmt"
	"strings"
)
//402. 移掉K位数字
//单调栈、贪心
func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	//单个字母的情况
	stack = stack[:len(stack)-k]
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		return "0"
	}
	return res
}


func main()  {
	num := "1432219"
	k := 3
	fmt.Println(removeKdigits(num, k))
}
