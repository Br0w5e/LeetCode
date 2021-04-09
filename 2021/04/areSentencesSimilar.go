package main

//5706. 句子相似性 III

import "strings"

//模拟
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	sentence1Arr := strings.Split(sentence1, " ")
	sentence2Arr := strings.Split(sentence2, " ")
	m, n := len(sentence1Arr), len(sentence2Arr)
	//从左边开始
	l1, l2 := 0, 0
	for l1 < m && l2 < n && sentence1Arr[l1] == sentence2Arr[l2] {
		l1++
		l2++
	}

	//从右边开始
	r1, r2 := m-1, n-1
	for r1 >= 0 && r2 >= 0 && sentence1Arr[r1] == sentence2Arr[r2] {
		r1--
		r2--
	}
	return l1 > r1 || l2 > r2

}
