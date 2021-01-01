package main

//667. 优美的排列 II
func constructArray(n int, k int) []int {
	res := make([]int, n)
	tmp := 1
	for i := 0; i <= k; i += 2 {
		res[i] = tmp
		tmp++
	}
	tmp = k+1
	for i := 1; i <= k; i += 2 {
		res[i] = tmp
		tmp--
	}

	for i := k+1; i < n; i++ {
		res[i] = i+1
	}
	return res
}
