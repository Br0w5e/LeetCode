//计算小于等于n的数字中共有多少个1
//暴力算法：超时
func countDigitOne(n int) int {
    ret := 0
    for i := 0; i <= n; i++{
        ret += conutOne(i)
    }
    return ret
}

func conutOne(n int) int {
    res := make([]int, 0)
    for n != 0 {
        res = append(res, n%10)
        n /= 10
    }
    ret := 0
    for _, i := range res {
        if i == 1{
            ret++
        }
    }
    return ret
}

func countDigitOne(n int) int {
	cnt := 0
	for i := 1; i <= n; i *= 10{
		divider := i * 10
		cnt +=(n / divider) * i + min(max(n%divider-i+1, 0), i)
	}
	return cnt
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}