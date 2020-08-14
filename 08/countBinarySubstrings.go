package main

import "fmt"

// 计算具有相同0和1的字串个数
// 方法
func countBinarySubstrings(s string) int {
	counts := make([]int, 0)
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
	}
	ans := 0
	for i := 1; i < len(counts); i++ {
		ans += min(counts[i], counts[i-1])
	}
	return ans
}

// 优化

func countBinarySubstrings1(s string) int {
	ptr, last, res := 0, 0, 0
	n := len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		res += min(count, last)
		last = count
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main()  {
	s := "00110"
	fmt.Println(countBinarySubstrings(s))
}