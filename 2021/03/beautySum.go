package main

//1781. 所有子字符串美丽值之和
import "fmt"

//map暴力--->超时
func beautySum(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i + 3; j <= len(s); j++ {
			res += beauty(s[i:j])
		}
	}
	return res
}

func beauty(s string) int {
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}
	max, min := 0, len(s)
	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

//数组暴力-->通过
func beautySum(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i + 3; j <= len(s); j++ {
			res += beauty(s[i:j])
		}
	}
	return res
}

func beauty(s string) int {
	m := make([]int, 26)
	for _, ch := range s {
		m[int(ch-'a')]++
	}
	max, min := 0, len(s)
	for _, v := range m {
		if v > max {
			max = v
		}
		if v > 0 && v < min {
			min = v
		}
	}
	return max - min
}

func main() {
	s := "aabcb"
	fmt.Println(beautySum(s))
}
