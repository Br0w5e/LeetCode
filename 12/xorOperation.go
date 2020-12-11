package main
//1486. 数组异或操作

func xorOperation(n int, start int) int {
	var xor int = 0
	for i:=0;i<n;i++ {
		xor ^= start + 2*i
	}
	return xor
}

func xorOperation2(n int, start int) int {
	if (start & 3) < 2 {
		if (n & 1) == 0 {
			return n & 3
		} else {
			return start + 2 * n - 3 + (n & 3)
		}
	} else {
		if (n & 1) == 0 {
			return (start + (n - 1) * 2) ^ (start - 2 + (n & 3))
		} else {
			return start + 1 - (n & 3)
		}
	}
}
