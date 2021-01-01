package main

//将奇数数列每次选取两个数进行加一和减一操作知道所有数字想的
func minOperations(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(2*i+1-n)
	}
	return sum/2
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}