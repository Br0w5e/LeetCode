package main

//319. 灯泡开关
import (
	"fmt"
	"math"
)

//一个数的因子要么是奇数个，要么是偶数个。  当因子为偶数的时候该数字肯定为完全平方数
//比n小的完全平方数只有根号n个
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

//暴力模拟--->超时
func bulbSwitch1(n int) int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		for j := i-1; j < n; j = j+i{
			nums[j] ^= 1
		}
		fmt.Println(nums)
	}
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

func main()  {
	bulbSwitch1(3)
}