package main

//326. 3的幂

//一直除
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}
