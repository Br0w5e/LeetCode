package main

import "strings"

func patternMatching(pattern string, value string) bool {
	if len(pattern) == 0 {
		//pattern == "" && value == ""
		if value == "" {
			return true
		}
		//pattern == "" && value != ""
		return false
	}
	if value == "" {
		//value == "" && pattern 包括a和b
		if strings.Contains(pattern, "a") && strings.Contains(pattern, "b") {
			return false
		}
		//value == "" && pattern 包括a或b
		return true
	}
	first := 0
	//与首字母相同字母个数
	for i, _ := range pattern {
		if pattern[i] == pattern[0] {
			first++
		} else {
			break
		}
	}
	//首字母个数和第二个字母个数
	firstAll, secondAll := 0, 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == pattern[0] {
			firstAll++
		} else  {
			secondAll++
		}
	}

	for i := 0; i <= len(value); i++ {
		a := value[:i]
		for j := first*i; j <= len(value); j++ {
			b := value[first*i: j]
			if len(value) != firstAll*len(a)+secondAll*len(b) {
				continue
			}
			var tmp string
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == pattern[0] {
					tmp += a
				} else {
					tmp += b
				}
				if value[:len(tmp)] != tmp {
					continue
				}
			}
			if value == tmp && a != b {
				return true
			}
		}
	}
	return false
}

//方法二
func patternMatching(pattern string, value string) bool {
	countA, countB := 0, 0
	//统计a和b数量
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			countA++
		} else {
			countB++
		}
	}
	//保持countA为多数
	if countA < countB {
		countA, countB = countB, countA
		tmp := ""
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'a' {
				tmp += "b"
			} else {
				tmp += "a"
			}
		}
		pattern = tmp
	}
	if len(value) == 0 {
		return countB == 0
	}
	if len(pattern) == 0 {
		return false
	}

	for lenA := 0; countA * lenA <= len(value); lenA++ {
		rest := len(value) - countA * lenA
		if (countB == 0 && rest == 0) || (countB != 0 && rest % countB == 0) {
			var lenB int
			if countB == 0 {
				lenB = 0
			} else {
				lenB = rest / countB
			}
			pos, correct := 0, true
			var valueA, valueB string
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == 'a' {
					sub := value[pos:pos+lenA]
					if len(valueA) == 0 {
						valueA = sub
					} else if valueA != sub {
						correct = false
						break
					}
					pos += lenA
				} else {
					sub := value[pos:pos+lenB]
					if len(valueB) == 0 {
						valueB = sub
					} else if valueB != sub {
						correct = false
						break
					}
					pos += lenB
				}
			}
			if correct && valueA != valueB {
				return true
			}
		}
	}
	return false
}