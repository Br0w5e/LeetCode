package main

//1578. 避免重复字母的最小删除成本
//last 记录当前进度， 如果删除last 则需要令last = i
func minCost(s string, cost []int) int {
	if len(s) < 2 {
		return 0
	}
	res := 0
	last := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[last] {
			if cost[last] < cost[i] {
				res += cost[last]
				last = i
			} else {
				res += cost[i]
			}
			continue
		}
		last = i
	}
	return res
}
