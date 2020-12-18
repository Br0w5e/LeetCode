package main

//5625. 比赛中的配对次数

func numberOfMatches(n int) int {
	res := 0
	for n != 1 {
		res += n/2
		if n%2 == 1 {
			n = (n-1)/2 + 1
		} else {
			n /= 2
		}
	}
	return res
}

func numberOfMatches1(n int) int {
	return n-1
}