package main
//暴力模拟-->超时
func maxScoreSightseeingPair(A []int) int {
	max := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i+1; j < len(A); j++ {
			if A[i]+A[j]+i-j > max {
				max = A[i]+A[j]+i-j
			}
		}
	}
	return max
}

//优化
func maxScoreSightseeingPair1(A []int) int {
	res, mx := 0, A[0] + 0 //mx存放和（A[i]+i），res存放结果（A[j]-j+mx)
	for j := 1; j < len(A); j++ {
		if mx+A[j]-j > res {
			res = mx+A[j]-j
		}
		if A[j]+j > mx {
			mx = A[j]+j
		}
	}
	return res
}