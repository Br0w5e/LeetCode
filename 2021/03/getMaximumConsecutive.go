package main

//5712. 你能构造出连续值的最大数目
import "sort"

//动态规划
//[0-x], [x, x+y] 就看y在不在了
func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	if coins[0] != 1 {
		return 1
	}
	res := 1
	for i := 1; i < len(coins); i++ {
		if coins[i] <= res+1 {
			res += coins[i]
		} else {
			break
		}
	}
	return res+1
}
