package main

import "sort"

//47. 全排列 II

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}

	m := make([]bool, len(nums))

	sort.Ints(nums)
	dfs(nums, nil, m, &res)
	return res
}

func dfs(nums []int, path []int, m []bool, res *[][]int) {
	//递归出口
	if len(path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if m[i] { //该数被使用
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && m[i-1] == false {
			continue
		}
		m[i] = true
		dfs(nums, append(path, nums[i]), m, res)
		m[i] = false
	}

}
