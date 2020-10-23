package main

//844. 比较含退格的字符串
func backspaceCompare(S string, T string) bool {
	return build(S) == build(T)
}

func build(s string) string {
	res := make([]byte, 0)
	for i, ch := range s {
		if ch != '#' {
			res = append(res, s[i])
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}

func backspaceCompare1(S string, T string) bool {
	ps, pt := len(S)-1, len(T)-1
	backspaceS, backspaceT := 0, 0
	for ps >= 0 || pt >= 0 {
		for ps >= 0 {
			if S[ps] == '#' {
				backspaceS++
				ps--
			} else if backspaceS > 0 {
				backspaceS--
				ps--
			} else {
				break
			}
		}
		for pt >= 0 {
			if T[pt] == '#' {
				backspaceT++
				pt--
			} else if backspaceT > 0 {
				backspaceT--
				pt--
			} else {
				break
			}
		}
		if ps >= 0 && pt >= 0 {
			if S[ps] != T[pt] {
				return false
			}
		}else if ps >= 0 || pt >= 0{
			return false
		}
		//相同字符
		ps--
		pt--
	}
	return true
}