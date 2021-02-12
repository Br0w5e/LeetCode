package main

//665. 非递减数列

//非递减数列
func checkPossibility(nums []int) bool {
	flag := 0
	for i := 1; i < len(nums) && flag < 2; i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		flag++
		if i-2 >= 0 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[1]
		}
	}
	return flag < 2
}