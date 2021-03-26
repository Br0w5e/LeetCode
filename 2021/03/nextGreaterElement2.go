package main

//556. 下一个更大元素 III
import (
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	digits := make([]int, 0)
	//取出各位数字，从低位到高位存储
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	//只有一位的情况
	if len(digits) <= 1 {
		return -1
	}

	bp := 0

	//从低位到高位寻找一个满足条件的数字：高位数字比低位小
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			bp = i
			break
		}
	}

	//没有找到这样的数字，说明该数字本来就是最大的
	if bp == 0 {
		return -1
	}

	//寻找低位第一个比bp点上的数字小的位置
	for i := 0; i < bp; i++ {
		if digits[i] > digits[bp] {
			digits[i], digits[bp] = digits[bp], digits[i]
			//逐个交换
			for j := 0; j < bp/2; j++ {
				digits[j], digits[bp-j-1] = digits[bp-j-1], digits[j]
			}
			break
		}
	}

	b := make([]byte, len(digits))
	for i := 0; i < len(digits); i++ {
		b[len(digits)-1-i] = byte('0' + digits[i])
	}

	s := string(b)

	next, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	if next > math.MaxInt32 {
		return -1
	}

	return next
}