package main

//1374. 生成每种字符都是奇数个的字符串
func generateTheString(n int) string {
	res := ""
	if n % 2 == 1 {
		for n > 0 {
			res += "a"
			n--
		}
	} else {
		for n > 1 {
			res += "a"
			n--
		}
		n--
		res += "b"
	}
	return res
}
