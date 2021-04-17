package main

func numSub(s string) int {
	res := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			sum++
		} else {
			sum = 0
		}
		res += sum
	}
	return res%(1000000000+7)
}
