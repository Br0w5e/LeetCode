package main

//1328. 破坏回文串

//模拟就行
func breakPalindrome(palindrome string) string {
	//小于1个无法操作
	if len(palindrome) <= 1 {
		return ""
	}
	//正常情况下，一般情况遇见大于‘a'替换成'a'即可
	s := []byte(palindrome)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] > 'a' {
			s[left] = 'a'
			return string(s)
		}
		left++
		right--
	}
	//全是a最后一个替换成b就行了
	s[len(s)-1]='b'
	return string(s)
}