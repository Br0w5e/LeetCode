package main

//1576. 替换所有的问号
func modifyString(s string) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			res[i] = s[i]
			continue
		}
		for j := byte('a'); j <= 'z'; j++ {
			if (i == 0 || j != res[i-1]) && (i == len(s)-1 || j != s[i+1]) {
				res[i] = j
				break
			}
		}
	}
	return string(res)
}
