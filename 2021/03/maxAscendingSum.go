package main

//5709. 最大升序子数组和

//模拟
func maxAscendingSum(nums []int) int {
	res := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			if sum > res {
				res = sum
			}
			sum = nums[i]
		}
	}
	if sum > res {
		return sum
	}
	return res
}
