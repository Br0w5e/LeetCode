package main


//456. 132模式
import (
	"fmt"
	"math"
)

//单调栈
func find132pattern(nums []int) bool {
	mid := math.MinInt64
	stack := make([]int, 0)
	for i := len(nums)-1; i >= 0; i-- {
		if mid > nums[i] {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			mid = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main()  {
	nums := []int{3, 1, 4, 2}
	fmt.Println(find132pattern(nums))
}
