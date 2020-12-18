package main

//438. 找到字符串中所有字母异位词

//滑动窗口
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	m := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		//p的贡献是+1， s的贡献是-1
		m[p[i]]++
		m[s[i]]--
	}

	for i := 0; i < len(s)-len(p)+1; i++ {
		if i > 0 {
			//s[i-1]出窗口，则需要+1
			m[s[i-1]]++
			m[s[i+len(p)-1]]--
		}
		//标记
		flag := true
		for _, v := range m {
			if v != 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return res
}
