package main

import "fmt"

//1218. 最长定差子序列
func longestSubsequence(arr []int, difference int) int {
	//存储以数字k结尾的长度v，等差数列
	m := make(map[int]int)
	for _, num := range arr {
		m[num] = m[num-difference]+1
	}
	max := 1
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}


func main()  {
	arr := []int{1,5,7,8,5,3,4,2,1}
	difference := -2
	fmt.Println(longestSubsequence(arr, difference))
}