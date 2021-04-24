package main

//1033. 移动石子直到连续
//模拟

func numMovesStones(a int, b int, c int) []int {
	res := make([]int, 2)
	maxNum := max(max(a, b), c)
	minNum := min(min(a, b), c)
	midNum := a+b+c-maxNum-minNum
	if maxNum-minNum+1 - 3 > 0 {
		res[1] = maxNum-minNum+1 - 3
		if maxNum - midNum== +1 || midNum - minNum == + 1 || midNum - minNum == 2 || maxNum - midNum == 2{
			res[0] = 1
		} else {
			res[0] = 2
		}
	} else {
		res[1] = 0
		res[0] = 0
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}