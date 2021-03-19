package main

//5701. 仅执行一次字符串交换能否使两个字符串相等
func areAlmostEqual(s1 string, s2 string) bool {
	//判断字母种类是否相同
	m := make(map[rune]int)
	for _, ch := range s1 {
		m[ch]++
	}
	for _, ch := range s2 {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}
	//判断需要异位的次数
	res := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			res++
		}
	}
	return res <= 2
}
