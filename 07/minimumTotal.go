package main

import "math"

//未优化
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	//本题中的相邻节点不包含上层中坐标比该节点大的节点
	//构建dp数组
	for i := 1; i < n; i++ {
		//最左边节点
		dp[i][0] = dp[i-1][0]+triangle[i][0]
		//一般节点
		for j := 1; j < i; j++{
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j])+triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1]+triangle[i][i]
	}
	//寻找dp数组中最后一行最小的
	left, right := 0, n-1
	for left < right {
		if dp[n-1][left] < dp[n-1][right] {
			right--
		} else {
			left++
		}
	}
	return dp[n-1][left]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//由于dp状态只与前一状态有关，跟更古老转态无关，因此我们可以进行空间优化
//底部开始
func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)
	for i := n-1; i >= 0 ; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1])+triangle[i][j]
		}
	}
	return dp[0]
}

func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j - 1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	left, right := 0, n-1
	for left < right {
		if dp[left] < dp[right] {
			right--
		} else {
			left++
		}
	}
	return dp[left]
}