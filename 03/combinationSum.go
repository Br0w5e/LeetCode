//数组组合：找出候选数组中所有可以组成目标数组的数字次数不要求
//dfs + 剪枝
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

// 数组组合：找出候选数组中所有可以组成目标数组的数字次数仅出现一次
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