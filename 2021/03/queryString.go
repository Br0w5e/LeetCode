package main

//1016. 子串能表示从 1 到 N 数字的二进制串
import (
	"fmt"
	"strings"
)

//将每个数字转换为二进制，并使用Contains进行判断即可
func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		num := fmt.Sprintf("%b", i)
		if !strings.Contains(S, num) {
			return false
		}
	}
	return true
}