package main

//338. 比特位计数

//比特位计数，计算每一个小于等于k的每个数字中有多少个1
//方法1：朴素方法
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = getOne(i)
	}
	return res
}

func getOne(num int) int {
	res := 0
	for num != 0 {
		res += num % 2
		num /= 2
	}
	return res
}

//方法2：i
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		dp[i] = dp[i>>1] + i%2
	}
	return dp
}

//方法3
func countBits(num int) []int {
    dp := make([]int, num+1)
    for i := 1; i < num+1; i++ {
        dp[i] = dp[i&(i-1)]+1
    }
    return dp
}

