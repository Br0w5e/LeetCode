package main

//1024. 视频拼接

//dp[i]表示将区间[0,i)覆盖所需的最少子区间的数量
//(l, i) (i, r)
//dp[i] = min(dp[l]+1, dp[i])
func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	for i := 0; i < T+1; i++ {
		//数据范围
		dp[i] = 101
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for _, clip := range clips {
			l, r := clip[0], clip[1]
			if l < i && i <= r {
				dp[i] = min(dp[l]+1, dp[i])
			}
		}
	}
	if dp[T] == 101 {
		return -1
	}
	return dp[T]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//贪心算法
//转化为第55. 跳跃游戏 canJump
func videoStitching2(clips [][]int, T int) int {
	res := 0
	m := make([]int, T)
	for _, clip := range clips {
		if clip[0] < T {
			m[clip[0]] = max(m[clip[0]], clip[1])
		}
	}
	pre, last := 0, 0
	for k, v := range m {
		if v > last {
			last = v
		}
		//最后一个还没到达
		if k == last {
			return -1
		}
		if k == pre {
			res++
			pre = last
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}