package main

import (
	"sort"
	"strings"
)

//242. 有效的字母异位词
//字母易位次
//两个字符串
func isAnagram(s string, t string) bool {
	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	sort.Strings(sArr)
	sort.Strings(tArr)
	return strings.Join(sArr, "") == strings.Join(tArr, "")
}

//49. 字母异位词分组
//字符串数组中的易位次
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	if len(strs) == 0 {
		return res
	}
	m := make(map[string][]string)
	for _, str := range strs {
		strArr := strings.Split(str, "")
		sort.Strings(strArr)
		index := strings.Join(strArr, "")
		m[index] = append(m[index], str)
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
