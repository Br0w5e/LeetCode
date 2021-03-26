package main

//532. 数组中的 k-diff 数对

//考虑到有重复的，我们使用map来进行计数，每次存放他们的和，最终返回map的长度即可
func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				//如果和相等，且差相等则这两数相等
				//也可以选择较小或者较大值存储
				m[nums[i]+nums[j]]++
				//存储较大值，用来去重
				//if nums[i] > nums[j] {
				//	m[nums[i]]++
				//} else {
				//	m[nums[j]]++
				//}
			}
		}
	}
	return len(m)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}