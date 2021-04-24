package main

//1744. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
//不太会
func canEat(candiesCount []int, queries [][]int) []bool {
	pre := make([]int, len(candiesCount)+1)
	for i := 1; i <= len(candiesCount); i++ {
		pre[i] = pre[i-1] + candiesCount[i-1]
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		//
		sum := pre[q[0]+1]
		if sum <= q[1] {
			res[i] = false
			continue
		}
		if q[2]*(q[1]+1) <= pre[q[0]] {
			res[i] = false
			continue
		}
		res[i] = true
	}
	return res
}
