//39 数组组合：找出候选数组中所有可以组成目标数组的数字次数不要求
//dfs + 剪枝
package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res)
	return res
}

func dfs(candidates []int, nums []int, target int, left int, res *[][]int) {
	//递归出口
	if target == 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := left; i < len(candidates); i++ {
		//剪枝
		if target < candidates[i] {
			return 
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res)
	}
}

// 40 数组组合：找出候选数组中所有可以组成目标数组的数字次数仅出现一次
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates) //快排
    res := [][]int{}
    dfs(candidates, nil, target, 0, &res)
    return res
}

func dfs(candidates []int, nums []int, target int, left int, res *[][]int) {
    if target == 0 { //结算
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        *res = append(*res, tmp)
        return 
    }
    for i := left; i < len(candidates); i++{
        if i != left && candidates[i] == candidates[i-1] {  //同层节点，数值相同 剪枝
            continue
        }
        if target < candidates[i] {//剪枝
            return
        }
        dfs(candidates, append(nums, candidates[i]), target-candidates[i], i+1, res)
    }
}
// 216 组合总和 III 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
func combinationSum3(k int, n int) [][]int {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := [][]int{}
	dfs(candidates, nil, n, 0, &res, k)
	return res
}

func dfs(candidates []int, nums []int, n int, left int, res *[][]int, k int) {
	if n == 0 && len(nums) == k{ //结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	if n == 0 || len(nums) > k{ //虽然满足但是长度不想等
		return
	}
	for i := left; i < len(candidates); i++{
		if i != left && candidates[i] == candidates[i-1] {  //同层节点，数值相同 剪枝
			continue
		}
		if n < candidates[i] {//剪枝
			return
		}
		dfs(candidates, append(nums, candidates[i]), n-candidates[i], i+1, res, k)
	}
}

//377 可以组成数字的个数
//超时
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums) //快排
	res := 0
	dfs(nums, target, &res)
	return res
}

func dfs(nums []int, target int, res *int) {
	//递归出口
	if target == 0 {
		*res++
		return
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(nums); i++ {
		dfs(nums, target-nums[i], res)
	}
}
//dp解决4问题
func combinationSum42(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	if target < nums[0] {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}