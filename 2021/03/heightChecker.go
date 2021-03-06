package main

//1051. 高度检查器
//生成同等长度的数组
import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	nums := make([]int, len(heights))
	copy(nums, heights)
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		if heights[i] != nums[i] {
			res++
		}
	}
	return res
}

func main()  {
	heights := []int{1,1,4,2,1,3}
	fmt.Println(heightChecker(heights))
}