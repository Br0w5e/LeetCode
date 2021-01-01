package main

//将0移到末尾
func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			nums[i-count], nums[i] = nums[i], nums[i-count]
		}
	}
}

func moveZeroes2(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
