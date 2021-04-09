package main

import "math"

//暴力---->超时
func countPairs(deliciousness []int) int {
	res := 0
	for i := 0; i < len(deliciousness); i++ {
		for j := i+1; j < len(deliciousness); j++ {
			if isPowerOfTwo(deliciousness[i]+deliciousness[j]) {
				res++
			}
		}
	}
	return res%1000000007
}

func isPowerOfTwo(n int) bool {
	//2的幂和小他1的就是0
	return n > 0 && (n & (n-1)) == 0
}


//map标记
func countPairs(deliciousness []int) int {
	cnt := 0
	freq := make(map[int]int)

	twoRet := make(map[int]int)
	twoRet[0] = 1
	twoItem := 1
	for i := 1; i <= 21; i++ {
		twoItem = 2 * twoItem
		twoRet[i] = twoItem
	}

	for _, val := range deliciousness {
		for i := 0; i <= 21; i++ {

			cnt += (freq[twoRet[i] - val])
			cnt =  cnt % (int(math.Pow10(9))+7)

		}
		freq[val] ++
	}

	return cnt
}