package main

//441. 排列硬币
func arrangeCoins(n int) int {
	//前面和
	sum := 0
	for i := 1; ;i++ {
		if sum + i > n {
			return i-1
		} else if sum + i == n {
			return i
		} else {
			sum += i
		}
	}
	return -1
}
