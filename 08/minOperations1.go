package main

import (
	"fmt"
	"math"
)

func minOperations(nums []int) int {
	res := 0
	for i := 0; i < math.MaxInt32; i++ {
		if contains_odd(nums) {
			res += do_odd(nums)
		}
		if all_zero(nums) {
			return res
		}
		res += do_ot(nums)
	}
	return res
}
func contains_odd(nums []int) bool {
	for _, num := range nums {
		if num%2 == 1 {
			return true
		}
	}
	return false
}
func do_odd(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			res++
			nums[i]--
		}
	}
	return res
}

func do_ot(nums []int) int {
	for i := 0; i < len(nums); i++ {
		nums[i] /= 2
	}
	return 1
}

func all_zero(nums []int) bool {
	for _, num := range nums {
		if num > 0 {
			return false
		}
	}
	return true
}


func main()  {
	nums := []int{2, 2}
	fmt.Println(minOperations(nums))
}