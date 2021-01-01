package main

//290. 单词规律

//map
func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")
	patternArr := strings.Split(pattern, "")
	if len(sArr) != len(patternArr) {
		return false
	}
	m := make(map[string]string)
	for i := 0; i < len(sArr); i++ {
		if val, ok := m[patternArr[i]]; ok {
			if val != sArr[i] {
				return false
			}
		} else {
			for _, v := range m {
				if v == sArr[i] {
					return false
				}
			}
			m[patternArr[i]] = sArr[i]
		}
	}
	return true
}
