//最小差值I
//没读懂它
func smallestRangeI(A []int, K int) int {
	max := A[0]
	min := A[0]
	for _, v := range A {
		if max < v {
			max = v
		}

		if min > v {
			min = v
		}
	}

	ret := max - min - 2*K
	if ret < 0 {
		ret = 0
	}
	return ret
}

func smallestRangeI(A []int, K int) int {
    minNum, maxNum := A[0], A[0]
    for _, num := range A {
        minNum = min(minNum, num)
        maxNum = max(maxNum, num)
    }
    return max(0, maxNum-minNum-2*K)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}