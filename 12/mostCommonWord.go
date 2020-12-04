package main

func mostCommonWord(paragraph string, banned []string) string {
	//map表生成
	m := make(map[string]int)
	str := ""
	for i := 0; i < len(paragraph); i++ {
		if paragraph[i] >= 'A' && paragraph[i] <= 'Z' {
			str += string(paragraph[i] + 32)
		} else if paragraph[i] >= 'a' && paragraph[i] <= 'z' {
			str += string(paragraph[i])
		} else if str != "" {
			m[str]++
			str = ""
		}
	}
	//最后一个，很关键
	if str != "" {
		m[str]++
	}
	//出现在banned表中的
	for i := 0; i < len(banned); i++ {
		m[banned[i]] = 0
	}
	//寻找结果
	res := ""
	max := 0
	for k, v := range m {
		if v > max {
			res = k
			max = v
		}
	}
	return res
}
