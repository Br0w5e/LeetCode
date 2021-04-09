package main

import "fmt"

//1209. 删除字符串中的所有相邻重复项 II
func removeDuplicates(s string, k int) string {
	bytes := []byte(s)
	counter := make([]int, len(s))
	for i := 0; i < len(bytes); i++ {
		//前后不相同，记录个数为1
		if i == 0 || bytes[i] != bytes[i-1] {
			counter[i] = 1
		} else {
			//前后相同，记录数+1
			counter[i] = counter[i-1] + 1
			//等于k，前移k个
			if counter[i] == k {
				bytes = append(bytes[:i-k+1], bytes[i+1:]...)
				i -= k
			}
		}
	}
	return string(bytes)
}


//直接操作字符串，字符串不能改变单个字母的内容，但是可以拼接
func removeDuplicates(s string, k int) string {
	counter := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] != s[i-1] {
			counter[i] = 1
		} else {
			counter[i] = counter[i-1]+1
			if counter[i] == k {
				s = s[:i-k+1]+s[i+1:]
				i -= k
			}
		}
	}
	return s
}

func main()  {
	s := "deeedbbcccbdaa"
	k := 3
	fmt.Printf("%s", removeDuplicates(s, k))
}