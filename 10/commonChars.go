package main

//1002. 查找常用字符
func commonChars(A []string) []string {
	//记录最小的出现次数
	m := make(map[rune]int)
	for _, ch := range A[0] {
		m[ch]++
	}
	for _, s := range A {
		//记录s中字母出现的次数
		m1 := make(map[rune]int)
		for _, ch := range s {
			m1[ch]++
		}
		//更改m
		for k, v := range m {
			m[k] = min(v, m1[k])
		}
	}
	//最终结果k
	res := make([]string, 0)
	for k, v := range m {
		for v > 0 {
			res = append(res, string(k))
			v--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
