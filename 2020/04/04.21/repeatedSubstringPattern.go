//判断字符串是不是重复字符串
func repeatedSubstringPattern(s string) bool {
	str1 := s + s
    str2 := str1[1 : len(str1)-1]
	if strings.Contains(str2,s) {
		return true
	}else {
		return false
	}
}