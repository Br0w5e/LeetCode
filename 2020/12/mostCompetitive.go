package main

//5614. 找出最具竞争力的子序列

//逐个选择
func mostCompetitive(nums []int, k int) []int {
	//非递减
	if judge(nums) {
		res := make([]int, k)
		copy(res, nums[:k])
		return res
	}

	res := make([]int, k)
	n := len(nums)
	start := 0
	cnt := 0
	for k > 0 {
		id := getSmall(nums, start, n-k)
		res[cnt] = nums[id]
		start = id + 1
		k--
		cnt++
	}
	return res
}

func getSmall(nums []int, start int, end int) int {
	for start < end {
		if nums[start] <= nums[end] {
			end--
		} else {
			start++
		}
	}
	return start
}

func judge(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}