package main

//1461. 检查一个字符串是否包含所有长度为 K 的二进制子串

//使用数组记录，当对应下标的数字不出现时，则返回false
import (
	"fmt"
)

func hasAllCodes(s string, k int) bool {
	nums := make([]int, 1<<k)
	for i := 0; i < len(s)-k+1; i++ {
		num := 0
		for j := i + k - 1; j >= i; j-- {
			num = num*2 + int(s[j]-'0')
		}
		nums[num] = 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 1 {
			return false
		}
	}
	return true
}

func main() {
	s := "00110110"
	k := 2
	fmt.Println(hasAllCodes(s, k))
}
