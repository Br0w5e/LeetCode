//字符串的最大公约数
func gcdOfStrings(str1 string, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)
	if str1+str2 != str2+str1{
		return ""
	}
	n := gcd(n1, n2)
	return str1[:n]
}

func gcd(n1, n2 int) int{
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	for n1 != 0{
		n1, n2 = n2%n1, n1
	}
	return n2
}