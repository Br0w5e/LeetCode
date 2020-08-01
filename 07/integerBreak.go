package main

//数学方法
func integerBreak(n int) int {
	if n <= 3 {
		return n-1
	}
	a, b := n/3, n%3
	if b == 0 {
		//余数为0 直接3次幂
		return pow3N(a)
	} else if b == 1 {
		//余数为1 借一个3 然后 2*2
		return pow3N(a-1)*4
	}
	//余数为2
	return pow3N(a)*2
}

func pow3N(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 3
	}
	return res
}

//dp
func integerBreak2(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func integerBreak3(n int) int {
	if n < 4 {
		return n - 1
	}
	dp := make([]int, n + 1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = max(max(2 * (i - 2), 2 * dp[i - 2]), max(3 * (i - 3), 3 * dp[i - 3]))
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}