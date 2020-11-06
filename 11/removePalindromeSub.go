package main

//1332. 删除回文子序列

//如果不是回文串，第一次全部删除含有a的子序列，第二次删除含有b的子序列
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return 2
		}
	}
	return 1
}
