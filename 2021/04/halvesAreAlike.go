package main

//1704. 判断字符串的两半是否相似

//模拟
func halvesAreAlike(s string) bool {
	left, right := 0, 0
	for i := 0; i < len(s)/2; i++ {
		if judge(s[i]) {
			left++
		}
	}

	for i := len(s)/2; i < len(s); i++ {
		if judge(s[i]) {
			right++
		}
	}
	return left == right
}

func judge(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O','U':
		return true
	}
	return false
}
