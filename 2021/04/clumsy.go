package main

//1006. 笨阶乘
//找规律
func clumsy(N int) int {
	if N > 4 {
		switch N%4 {
		case 0:
			return N+1
		case 1, 2:
			return N+2
		case 3:
			return N-1
		}
	} else {
		switch N {
		case 0, 1, 2:
			return N
		case 3:
			return 6
		case 4:
			return 7
		}
	}
	return 0
}

//模拟
func clumsy(N int) int {
	if N <= 2 {
		return N
	}
	if N == 3 {
		return 6
	}
	sum := N*(N-1)/(N-2)+N-3
	N -= 4
	for N >= 4 {
		sum += (-N*(N-1))/(N-2)+N-3
		N -= 4
	}
	return sum - clumsy(N)
}