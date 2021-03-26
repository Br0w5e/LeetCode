package main


//456. 132模式
import (
	"fmt"
	"math"
)

//单调栈
func find132pattern(nums []int) bool {
	//中间大小的数字
	mid := math.MinInt64
	stack := make([]int, 0)
	//从后往前栈中存在的是2， 中间元素是3，
	for i := len(nums)-1; i >= 0; i-- {
		//当前元素小于中间元素
		if mid > nums[i] {
			return true
		}

		//单调递增栈，栈中存放的是最大数字 栈中的元素一定大于mid，由于是从后往前遍历的，所以该元素肯定在mid之前
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
