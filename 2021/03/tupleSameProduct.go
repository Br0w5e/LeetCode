package main

import "sort"

//二分超时
func tupleSameProduct(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := len(nums)-1; j >= i+3; j-- {
			mul1 := nums[i]*nums[j]
			left, right := i+1, j-1
			for left < right {
				mul2 := nums[left]*nums[right]
				if mul1 == mul2 {
					res += 2*2*2
					left++
					right--
				} else if mul1 < mul2 {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}


func tupleSameProduct1(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			m[nums[i]*nums[j]]++
		}
	}
	res := 0
	for _, v := range m {
		//内部交换排序总共有4种方式
		res += (v - 1) * v  * 4
	}
	return res
}
