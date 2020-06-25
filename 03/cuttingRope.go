//给定正整数求正整数最大分解的成绩
//剪绳子一
func pow3N(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 3
	}
	return res
}
func cuttingRope(n int) int {
    if n <= 3 {
        return n-1
    }
    a, b := n/3, n%3
    if b == 0{
        return pow3N(a)
    }
    if b == 1 {
        return pow3N(a-1)*4
    }
    return pow3N(a)*2
}

//剪绳子二
func pow3N(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res = (res * 3) % 1000000007
	}
	return res
}
func cuttingRope(n int) int {
    if n <= 3 {
        return n-1
    }
    a, b := n/3, n%3
    if b == 0{
        return pow3N(a)
    }
    if b == 1 {
        return pow3N(a-1)*4%1000000007
    }
    return pow3N(a)*2%1000000007
}
