package main

//1306. 跳跃游戏 III

//dfs
func canReach(arr []int, start int) bool {
	return dfs(arr, start)
}

func dfs(arr []int, st int) bool {
	if st < 0 || st >= len(arr) || arr[st] == -1 {
		return false
	}
	step := arr[st]
	arr[st] = -1
	return step == 0 || dfs(arr, st+step) || dfs(arr, st-step)
}
