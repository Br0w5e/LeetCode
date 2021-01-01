package main

//1480. 一维数组的动态和
//数组动态和

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1]+nums[i]
	}
	return res
}
