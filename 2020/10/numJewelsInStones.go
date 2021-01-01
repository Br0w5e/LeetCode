package main

//771. 宝石与石头

func numJewelsInStones(J string, S string) int {
	res := 0
	for i := 0; i < len(J); i++ {
		for j := 0; j < len(S); j++ {
			if J[i] == S[j] {
				res++
			}
		}
	}
	return res
}

func numJewelsInStones2(J string, S string) int {
	res := 0
	for _, j := range J {
		for _, s := range S {
			if j == s {
				res++
			}
		}
	}
	return res
}

func numJewelsInStones3(J string, S string) int {
	m := make(map[rune]int)
	for _, s := range S {
		m[s]++
	}
	res := 0
	for _, j := range J {
		res += m[j]
	}
	return res
}
