package main

//1324. 竖直打印单词
import "strings"

func printVertically(s string) []string {
	arr := strings.Split(s, " ")
	max := 0
	for _, a := range arr {
		if len(a) > max {
			max = len(a)
		}
	}
	res := make([]string, max)
	//竖直加入
	for i := 0; i < max; i++ {
		for j := 0; j < len(arr); j++ {
			if len(arr[j]) >= i+1 {
				res[i] += string(arr[j][i])
			} else {
				res[i] += " "
			}
		}
		//去除末尾空格
		for res[i][len(res[i])-1] == ' ' {
			res[i] = res[i][:len(res[i])-1]
		}
	}
	return res
}