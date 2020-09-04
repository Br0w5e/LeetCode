package main

import "fmt"

//205 同构字符串
func isIsomorphic(s string, t string) bool {
	mS, mT, n := [128]int{}, [128]int{}, len(s) // mS, mT分别记录s和t中每个字母的映射
	for i := 0; i < n; i++ {
		c1, c2 := s[i], t[i]
		if mS[c1] != mT[c2] { // c1和c2对应的映射不同时
			return false
		} else {
			if mS[c1] == 0 { // 当字母第一次出现时，将字符映射为所在位置加1
				mS[c1] = i+1
				mT[c2] = i+1
			}
		}
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1, m2 := make(map[rune]rune), make(map[rune]rune)

	for i, letter := range s {
		//m1映射
		if val, ok := m1[letter]; ok {
			if val != rune(t[i]) {
				return false
			}
		} else {
			m1[letter] = rune(t[i])
		}
		//m2映射
		if val, ok := m2[rune(t[i])]; ok {
			if val != letter {
				return false
			}
		} else {
			m2[rune(t[i])] = letter
		}
	}
	return true
}
func main()  {
	s := "egga"
	t := "adda"
	fmt.Println(isIsomorphic2(s, t))
}