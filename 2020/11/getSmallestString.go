package main

import "fmt"

//5606. 具有给定数值的最小字符串
//首先将所有都置为'a',然后从后往前遍历直到k为0
func getSmallestString(n int, k int) string {
	s := make([]byte, 0)
	for i := 0; i < n; i++ {
		s = append(s, 'a')
		k--
	}
	for i := n-1; i >= 0; i-- {
		num := min(k, 25)
		s[i] = byte(int(s[i])+num)
		k -= num
	}
	return string(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main()  {
	fmt.Println(getSmallestString(3, 27))
}