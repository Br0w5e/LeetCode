package main

import "sort"

//剪树
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 1)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	dfs(0, []int{}, &res, nums)
	return res
}

//回溯+剪枝
func dfs(start int, path []int, res *[][]int, nums []int)  {
	for i := start; i < len(nums); i++ {
		//只需要判断当前数字和上一个数字是否相同，相同的话跳过即可
		//同层剪纸
		//i > start 保证首次出现元素肯定会加入的
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		tmp := make([]int, len(path))
		copy(tmp, path)
		tmp = append(tmp, nums[i])
		*res = append(*res, tmp)
		dfs(i+1, tmp, res, nums)
	}
}


//因为输入数据已经排序， 在回溯时检查当前元素是否和上一个元素相同，即 nums[i - 1] == nums[i] 如果发现和上一个相同了就饿可以直接跳过
//需要注意的是 nums[i - 1] == nums[i] 的作用是在递归树的同一层上的， 也就是防止 nums = [1, 2, 2'] 出现 [1, 2]， [1, 2'] 两个重复的解 （即： 通过 nums[1] == nums[2]: continue 进行的剪枝）
//而解 [1, 2, 2'] 实际中的 2, 2' 是位于递归树的不同层级的， 也就是剪枝要从每一层的第二个元素开始， 所以还要增加一个 i > start 或者 i != start
//最终相对于子集 的思路就是要增加一个 i > start && nums[i - 1] == nums[i] 或 i != start && nums[i - 1] == nums[i] 的剪枝
