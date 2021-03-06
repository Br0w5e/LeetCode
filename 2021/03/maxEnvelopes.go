package main

//354. 俄罗斯套娃信封问题

//先将左端点按照递增的顺序排列（相同时右端点大的排在前面），再按照右端点求最长递增子序列
import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	//左端点按照递增排序，相同时将右端点大的排在前面
	sort.SliceStable(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	//最长递增子序列求法（以右端点为标准）
	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
