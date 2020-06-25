//反转给定的字符串
//方法：两个库函数的使用
//将www baidu com 转化为com baidu www
func reverseWords(s string) string {
	parts := strings.Split(s, " ")
	size := len(parts)
	var rparts[] string
	for i := size-1; i >= 0; i-- {
		if parts[i] != ""{
			rparts = append(rparts, parts[i])
		}
	}
	return strings.Join(rparts, " ")
}

//将www.baidu.com  转化为 com.baidu.www
func reverseWords(s string) string {
	parts := strings.Split(s, ".")
	size := len(parts)
	var rparts[] string
	for i := size-1; i >= 0; i-- {
		if parts[i] != "" {
			rparts = append(rparts, parts[i])
		}
	}
	return strings.Join(rparts, ".")
}

