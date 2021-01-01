package main

//5603. 确定两个字符串是否接近

import "sort"

func closeStrings(word1 string, word2 string) bool {
	m := len(word1)
	n := len(word2)
	if m != n {
		return false
	}
	//记录出现次数
	repeat1 := make([]int, 26)
	repeat2 := make([]int, 26)
	for i := 0; i < m; i++ {
		repeat1[int(word1[i]-'a')]++
		repeat2[int(word2[i]-'a')]++
	}
	//word1中出现而word2中没出现
	for i := 0; i < 26; i++ {
		if (repeat1[i] == 0 && repeat2[i] != 0) || (repeat1[i] != 0 && repeat2[i] == 0) {
			return false
		}
	}
	sort.Ints(repeat1)
	sort.Ints(repeat2)
	//排序比较出现次数
	for i := 0; i < 26; i++ {
		if repeat1[i] != repeat2[i] {
			return false
		}
	}
	return true
}