package main
//G(n) = G(0)*G(n-1)+G(1)*(n-2)+...+G(n-1)*G(0)
//dp
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}

//数学方法：卡塔兰数

func numTrees1(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res = res*2*(2*i+1)/(i+2)
	}
	return res
}
