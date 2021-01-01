package main

//334. 递增的三元子序列
import "math"

//逐个遍历，维护第二大和 最小元素， 当找到最大的时候直接返回
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min := math.MaxInt32
	max := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		} else if num > min && num <= max {
			max = num
		} else if num > max {
			return true
		}
	}
	return false
}
