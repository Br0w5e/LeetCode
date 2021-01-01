package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	left, right, res := make([]int, length), make([]int, length), make([]int, length)
	left[0] = 1
	for i := 1; i < length; i++ {
		left[i] = left[i-1]*nums[i-1]
	}
	right[length-1] = 1
	for i := length-2; i >= 0; i-- {
		right[i] = nums[i+1]*right[i+1]
	}
	for i := 0; i < length; i++ {
		res[i] = left[i]*right[i]
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	res[0] = 1
	for i := 1; i < length; i++ {
		res[i] = nums[i-1 ]*res[i-1]
	}
	r := 1
	for i := length-1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}
