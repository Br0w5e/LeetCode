package main


//暴力超时
func subarraysWithKDistinct(A []int, K int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		m := make(map[int]int, 0)
		for j := i; j < len(A); j++ {
			m[A[j]]++
			if len(m) == K {
				res++
			}
			if len(m) > K {
				continue
			}
		}
	}
	return res
}


//题目转换-->最多有K个减去最大有K-1个就是 恰好有K个
func subarraysWithKDistinct(A []int, K int) int {
	return atMostWithDistinctK(A, K) - atMostWithDistinctK(A, K-1)
}

func atMostWithDistinctK(A []int, K int) int {
	res := 0
	cnt := 0
	left, right := 0, 0
	freq := make(map[int]int, 0)
	for right < len(A) {
		if freq[A[right]] == 0 {
			cnt++
		}
		freq[A[right]]++
		right++
		//大于K的时候进行处理直到，cnt小于等于K
		for cnt > K {
			freq[A[left]]--
			if freq[A[left]] == 0{
				cnt--
			}
			left++
		}
		//小于K的时候结果
		res += right-left
	}
	return res
}
