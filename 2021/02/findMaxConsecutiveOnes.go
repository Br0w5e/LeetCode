package main
//485. 最大连续1的个数

//双指针
func findMaxConsecutiveOnes(nums []int) int {
	slow, fast := 0, 0
	res := 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			if fast-slow > res {
				res = fast-slow
			}
			slow = fast+1
		}
		fast++
	}
	if fast-slow > res {
		res = fast-slow
	}
	return res
}
