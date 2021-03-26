package main

//762. 二进制表示中质数个计算置位

//
func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		if judge(i) {
			res++
		}
	}
	return res
}

func judge(n int) bool  {
	ones := 0
	for n != 0 {
		ones += n%2
		n /= 2
	}
	switch ones {
	case 2, 3, 5, 7, 11, 13, 17, 19:
		return true
	}
	return false
}
