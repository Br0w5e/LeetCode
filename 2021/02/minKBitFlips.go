package main

//995. K 连续位的最小翻转次数

//暴力超时
func minKBitFlips(A []int, K int) int {
	res := 0
	for i := 0; i <= len(A)-K; i++ {
		if A[i] == 0 {
			reverse(A, K, i)
			res++
		}
	}
	for _, num := range A {
		if num == 0 {
			return -1
		}
	}
	return res
}

func reverse(A []int, K int, flag int) {
	for i := flag; i < flag+K; i++ {
		A[i] = 1-A[i]
	}
	return
}


func minKBitFlips(A []int, K int) int {
	n := len(A)
	diff := make([]int, n+1)
	cnt, res := 0, 0
	for i, v := range A {
		cnt += diff[i]
		//0翻转成功的cnt一定是奇数
		//1翻转成功的cnt一定是偶数
		if (v+cnt)%2 == 0 {
			if i+K > n {
				return -1
			}
			res++
			cnt++
			diff[i+K]--
		}
	}
	return res
}