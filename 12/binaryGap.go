package main

//868. 二进制间距
func binaryGap(n int) int {
	cnt := 0
	prev := -1
	res := 0
	for n != 0 {
		if n % 2 == 1 {
			if prev != -1 {
				res = max(res, cnt-prev)
			}
			prev = cnt
		}
		n /= 2
		cnt++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
