package main

import "fmt"

// 87. 扰乱字符串

//暴力判断--->超时
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	n := len(s1)
	m := make([]int, 26)
	//统计
	for i := 0; i < n; i++ {
		m[int(s1[i]-'a')]++
		m[int(s2[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if m[i] != 0 {
			return false
		}
	}
	//暴力搜索
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) || isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}

// dp--->通过
//  不会
func isScramble(s1, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	// 第一个字符串从 i1 开始，第二个字符串从 i2 开始，子串的长度为 length
	// 和谐返回 1，不和谐返回 0
	var dfs func(i1, i2, length int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()

		// 判断两个子串是否相等
		x, y := s1[i1:i1+length], s2[i2:i2+length]
		if x == y {
			return 1
		}

		// 判断是否存在字符 c 在两个子串中出现的次数不同
		freq := [26]int{}
		for i, ch := range x {
			freq[ch-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq[:] {
			if f != 0 {
				return 0
			}
		}

		// 枚举分割位置
		for i := 1; i < length; i++ {
			// 不交换的情况
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}
			// 交换的情况
			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}

		return 0
	}
	return dfs(0, 0, n) == 1
}

func main() {
	s1 := "great"
	s2 := "rgeat"
	fmt.Println(isScramble(s1, s2))
}
