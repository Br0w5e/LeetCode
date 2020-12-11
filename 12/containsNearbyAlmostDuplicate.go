package main

//220. 存在重复元素 III

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if j-i > k {
				break
			} else if abs(nums[i]-nums[j]) > t {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
