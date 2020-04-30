//完全平方和
//DP
func numSquares(n int) int {
	leng := int(math.Sqrt(float64(n)))+1 
	squareNums := make([]int, 0, leng)
	dp := make([]int, 1, n + 1)    
	for i:= 1; i<=leng;i++{
		squareNums = append(squareNums,i*i)    //获取平方数
	}
	for i:= 1;i<=n;i++{
		min := -1
		for idx := 0; squareNums[idx]<= i;idx ++{
			if min == -1 || dp[i-squareNums[idx]] < min{
				min = dp[i-squareNums[idx]]
			}
		}
		dp = append(dp,min + 1)
	}
	return dp[n]
}