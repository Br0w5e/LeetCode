package main

//5613. 最富有客户的资产总量
//模拟
func maximumWealth(accounts [][]int) int {
	res := 0
	for i := 0; i < len(accounts); i++ {
		tmp := 0
		for j := 0; j < len(accounts[0]); j++ {
			tmp += accounts[i][j]
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
