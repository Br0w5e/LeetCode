package main
//暴力
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s)+1; j++ {
			if judge(s[i:j]) {
				res++
			}
		}
	}
	return res
}

func judge(s string) bool {
	if len(s) == 1 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
//中心扩展
func countSubstrings1(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

//中心扩展2
func countSubstrings2(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += expand(i, i, s) //奇数扩展
		res += expand(i, i+1, s) //偶数扩展
	}
	return res
}

func expand(l, r int, s string) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		res++
	}
	return res
}