package main

import "fmt"

func minOperations(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d <= 0 {
			res += -d + 1
			nums[i] += -d + 1
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(minOperations(nums))
}
