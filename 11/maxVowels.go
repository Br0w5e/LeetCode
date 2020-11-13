package main

import "fmt"
//1456. 定长子串中元音的最大数目
//滑动窗口
func maxVowels(s string, k int) int {
	m := make(map[byte]bool)
	m['a'] = true
	m['e'] = true
	m['i'] = true
	m['o'] = true
	m['u'] = true

	tmp := 0
	for i := 0; i < k; i++ {
		if m[s[i]] {
			tmp++
		}
	}
	res := tmp
	for i := 0; i < len(s)-k; i++ {
		//出去的是元音，进来的不是
		if m[s[i]] && !m[s[i+k]] {
			tmp--
			//进来的是元音，出去的不是
		} else if !m[s[i]] && m[s[i+k]]{
			tmp++
		}
		//其他情况，进元音出元音，或者进辅音出辅音，不考虑

		if tmp > res{
			res = tmp
		}
	}
	return res
}

func main()  {
	s := "ibpbhixfiouhdljnjfflpapptrxgcomvnb"
	k := 33
	fmt.Println(maxVowels(s, k))
}
