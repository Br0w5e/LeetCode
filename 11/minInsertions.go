package main

//1541. 平衡括号字符串的最少插入次数
func minInsertions(s string) int {
	res, left := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			// 没有左括号了，需要加一个左括号，res++
			if left == 0 {
				res++
			} else {
				left--
			}
			// 以下两种情况只有一个右括号，需要再加一个右括号，res++
			if i == len(s)-1 || s[i+1] != ')' {
				res++
			} else {
				i++
			}
		}
	}
	return res + left*2
}
