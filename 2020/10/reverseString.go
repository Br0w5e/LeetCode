package main
//344. 反转字符串
//双指针
func reverseString(s []byte)  {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
	return
}
