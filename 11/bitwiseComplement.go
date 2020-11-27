package main

//1009. 十进制整数的反码

//模拟
func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	res := make([]int, 0)
	for N != 0 {
		res = append(res, N%2)
		N /= 2
	}
	for i := 0; i < len(res); i++ {
		res[i] ^= 1
	}
	ret := 0
	for i := len(res)-1; i >= 0; i-- {
		ret = ret*2 + res[i]
	}
	return ret
}


//反码加原码等于 1111111
func bitwiseComplement1(N int) int {
	if N == 0 {
		return 1
	}
	i := 1
	for i <= N {
		i <<= 1
	}
	//i-1 表示和N具有相同长度的全1数字
	return i-1 - N
}