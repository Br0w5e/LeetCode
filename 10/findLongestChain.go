package main

//646. 最长数对链
import "sort"

//dp
//dp[i]存储以pairs[i]结尾的最长链的长度
func findLongestChain(pairs [][]int) int {
	dp := make([]int, len(pairs))
	for i := 0; i < len(pairs); i++ {
		dp[i] = 1
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	res := 1
	//将数组按照起始节点递增的顺序排序
	for i := 1; i < len(pairs); i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(dp[i], res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//贪心
func findLongestChain2(pairs [][]int) int {
	//将数组按照起始节点递增的顺序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	res := 1
	tmp := pairs[0][1]

	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > tmp {
			res++
			tmp = pairs[i][1]
		}
	}
	return res
}