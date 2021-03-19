package main

//LCP 17. 速算机器人

//模拟
func calculate(s string) int {
	X, Y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			X = 2*X + Y
		} else {
			Y = 2*Y + X
		}
	}
	return X + Y
}
