package main

//5630. 删除子数组的最大得分

//滑动窗口
func maximumUniqueSubarray(nums []int) int {
	m := make(map[int]int)
	res, sum := 0, 0
	for i, j := 0, 0; i < len(nums); i++ {
		m[nums[i]]++
		sum += nums[i]
		for m[nums[i]] > 1 {
			m[nums[j]]--
			sum -= nums[j]
			j++
		}
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
