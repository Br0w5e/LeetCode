package main

//342. 4的幂

//移位，每次移2位
func isPowerOfFour(n int) bool {
	res := 1
	for res < n {
		res <<= 2
	}
	return res == n
}
