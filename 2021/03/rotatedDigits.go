package main

//788. 旋转数字

//每位都在[2, 5, 6, 9, 0, 1, 8]
//至少有一位在[2, 5, 6, 9]
func rotatedDigits(N int) int {
	res := 0
	for i := 1; i <= N; i++ {
		if judge(i) {
			res++
		}
	}
	return res
}

//每位都在[2, 5, 6, 9, 0, 1, 8]
func good1(n int) bool {
	return n == 0 || n == 1 || n == 8 || n == 2 || n == 5 || n == 6 || n == 9
}

//至少有一位在[2, 5, 6, 9]
func good2(n int) bool {
	return n == 2 || n == 5 || n == 6 || n == 9
}

func judge(n int) bool {
	tmp := n
	for tmp != 0 {
		if !good1(tmp % 10) {
			return false
		}
		tmp /= 10
	}
	for n != 0 {
		if good2(n % 10) {
			return true
		}
		n /= 10
	}
	return false
}

func rotatedDigits1(N int) int {
	res := 0
	for i := 2; i <= N; i++ {
		x := i
		flag := false
		for x > 0 {
			switch x % 10 {
			case 3, 4, 7:
				flag = false
				x = 0
			case 2, 5, 6, 9:
				flag = true
			}
			x = x / 10
		}
		if flag {
			res++
		}
	}
	return res
}
