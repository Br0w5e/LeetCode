package main
//回溯加剪枝
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	//cur当前下标，length路径长度，pre前一个数
	dfs := func(cur, length, pre int) {}
	dfs = func(cur, length, pre int) {
		//满足大于2的条件
		if length >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		//处理
		cache := make(map[int]int)
		//其他数字
		for i := cur; i < len(nums); i++ {
			_, ok := cache[nums[i]]
			//不需要处理
			if nums[i] < pre || ok {
				continue
			}
			path = append(path, nums[i])
			cache[nums[i]] = 1
			dfs(i+1, length+1, nums[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0,0, -101)
	return res
}
