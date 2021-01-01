package main
//按位相与
//求公共前缀的
func rangeBitwiseAnd(m, n int) int {
	shift := 0
	for m < n {
		m, n := m >> 1, n >> 1
		shift++
	}
	return m << shift
}

func rangeBitwiseAnd2(m, n int) int {
	for m < n {
		n &= (n-1)
	}
	return n
}
