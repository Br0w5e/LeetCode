package main

//85. 最大矩形
//参考84题
//dp[i][j] i,j为顶点的最大列连续矩
func maximalRectangle(matrix [][]byte) int {
	max := 0
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			//每列的第一个元素
			if i == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				//计算当前行的矩形的高度
				if matrix[i][j] == '1' {
					dp[i][j] = dp[i-1][j] + 1
				}
			}
		}
		t := largestRectangleArea(dp[i])
		if t > max {
			max = t
		}
	}

	return max
}

func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		minH := heights[i]
		for j := i; j < len(heights); j++ {
			if minH > heights[j] {
				minH = heights[j]
			}
			if max < minH*(j-i+1) {
				max = minH * (j - i + 1)
			}
		}
	}
	return max
}
