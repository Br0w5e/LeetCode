package main
//1318. 或运算的最小翻转次数
func minFlips(a int, b int, c int) int {
	res := 0
	for a | b != c {
		a1, b1, c1 := a&1, b&1, c&1
		a, b, c = a >> 1, b >> 1, c >> 1
		if a1 | b1 != c1 {
			//001、010、100、110，前面三种情况操作一次，最后一种操作两次
			if a1 == 1 && b1 == 1 && c1 == 0 {
				res += 2
			} else {
				res++
			}
		}
	}
	return res
}


