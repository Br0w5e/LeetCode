package main

import "sort"

//18. 四数之和
//暴力解法
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	if len(nums) < 4 {
		return nil
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		//去重
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i+1; j < len(nums)-2; j++ {
			//去重
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j+1; k < len(nums)-1; k++ {
				//去重
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				for l := k+1; l < len(nums); l++ {
					//去重
					if l != k+1 && nums[l] == nums[l-1] {
						continue
					}
					if nums[i] + nums[j] + nums[k] + nums[l] == target {
						tmp := make([]int, 4)
						tmp[0], tmp[1], tmp[2], tmp[3] = nums[i], nums[j], nums[k], nums[l]
						res = append(res, tmp)
					}
				}
			}
		}
	}
	return res
	bool
}


//小范围优化
func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	if n < 4 {
		return nil
	}
	res := make([][]int, 0)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		//去重,或者最大值小于
		if i != 0 && nums[i] == nums[i-1] || (nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target) {
			continue
		}
		for j := i+1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			//去重
			if j != i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for k := j+1; k < len(nums)-1 && nums[i]+nums[j]+nums[k]+nums[k+1] <= target; k++ {
				//去重
				if k != j+1 && nums[k] == nums[k-1] || nums[i] + nums[j] + nums[k] + nums[n-1] < target{
					continue
				}
				for l := k+1; l < len(nums) && nums[i]+nums[j]+nums[k]+nums[l] <= target; l++ {
					//去重
					if l != k+1 && nums[l] == nums[l-1] || nums[i] + nums[j] + nums[k] + nums[l] < target {
						continue
					}
					if nums[i] + nums[j] + nums[k] + nums[l] == target {
						tmp := make([]int, 4)
						tmp[0], tmp[1], tmp[2], tmp[3] = nums[i], nums[j], nums[k], nums[l]
						res = append(res, tmp)
					}
				}
			}
		}
	}
	return res
}