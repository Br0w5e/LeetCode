package main

//1523. 在区间范围内统计奇数数目

//两边都是偶数的情况，
//其他情况
func countOdds(low int, high int) int {
	if low%2 == 0 && high%2 == 0 {
		return (high-low)/2
	}
	return (high-low)/2 + 1
}