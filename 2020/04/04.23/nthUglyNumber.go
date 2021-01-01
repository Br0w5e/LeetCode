//丑数
func nthUglyNumber(n int) int {
    return nthUglyNumberN(n, 2, 3, 5)
}

func nthUglyNumberN(n int, a int, b int, c int) int {
	nth := 1
	uglyList := make([]int, n)
	uglyList[0] = 1

	idx_a := 0
	idx_b := 0
	idx_c := 0

	for {
		if nth == n {
			return uglyList[n-1]
		}

		xa := uglyList[idx_a] * a
		xb := uglyList[idx_b] * b
		xc := uglyList[idx_c] * c

		uglyList[nth] = min3(xa, xb, xc)

		if uglyList[nth] == xa {
			idx_a++
		}
		if uglyList[nth] == xb {
			idx_b++
		}
		if uglyList[nth] == xc {
			idx_c++
		}

		nth++
	}
}


func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func min2 (x, y int) int {
    if x < y {
        return x
    }
    return y
}