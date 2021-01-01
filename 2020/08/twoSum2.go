package main

import (
	"fmt"
	"math"
)

//计算骰子点数变化概率
//动态规划
func twoSum2(n int) []float64  {
	sum := math.Pow(6.0, float64(n))
	dp, res := make([]int, 6*n+1), []float64{}
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6*i; j >= i; j-- {
			dp[j] = 0
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[j] += dp[j-k]
				}
			}
		}
	}
	for k := n; k <= 6*n; k++ {
		res = append(res, float64(dp[k])/sum)
	}
	return res
}

func main()  {
	n := 3

	fmt.Println(twoSum2(n))
}
