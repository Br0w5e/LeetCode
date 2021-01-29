package main

import "unicode"

//1441. 用栈操作构建数组

//设计标志
func buildArray(target []int, n int) []string {
	res := make([]string, 0, n*2)
	flag := 1
	for _, num := range target {
		for flag != num {
			res = append(res, "Push", "Pop")
			flag++
		}
		res = append(res, "Push")
		flag++
	}
	return res
	unicode.IsUpper()
	unicode.IsLower()
}
