package main

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1)+climbStairs(n-2)
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 0; i < n; i++{
		a, b = b, a+b
	}
	return b
}

func main()  {
	println(climbStairs(4))
}