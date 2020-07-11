package main

func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int, 0)
	//k == 0 的情况
	if k == 0 {
		return res
	}
	// shorter == longer 的情况
	if shorter == longer {
		res = append(res, longer*k)
		return res
	}
	//一般情况
	for i := 0; i <= k ; i++ {
		res = append(res, i*longer+(k-i)*shorter)
	}
	return res
}
