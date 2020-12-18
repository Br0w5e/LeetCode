package main

//376. 摆动序列
//动态规划


//up[i] 表示以前i个元素中的某一个为结尾的最长的「上升摆动序列」的长度。
//down[i] 表示以前i个元素中的某一个为结尾的最长的「下降摆动序列」的长度。

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		up[i] = up[i-1]
		down[i] = down[i-1]
		if nums[i] > nums[i-1] {
			up[i] = max(up[i], down[i]+1)
		} else if nums[i] < nums[i-1] {
			down[i] = max(down[i], up[i]+1)
		}
	}
	return max(up[n-1], down[n-1])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}


func wiggleMaxLength1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = max(down, up+1)
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
