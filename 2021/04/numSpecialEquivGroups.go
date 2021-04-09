package main

import "sort"

//893. 特殊等价字符串组

//暴力模拟--->超时
func numSpecialEquivGroups(A []string) int {
	res := 0
	m := make(map[string]bool)
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if !m[A[i]] && judge(A[i], A[j]) {
				res++
				m[A[i]] = true
				m[A[j]] = true
			} else if judge(A[i], A[j]) {
				m[A[j]] = true
			}
		}
	}
	for _, a := range A {
		if m[a] == false {
			res++
		}
	}
	return res
}

func judge(s1 string, s2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			m1[s1[i]]++
			m1[s2[i]]--
		} else {
			m2[s1[i]]++
			m2[s2[i]]--
		}
	}
	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}

	for _, v := range m2 {
		if v != 0 {
			return false
		}
	}
	return true
}

//字符串按照奇欧排序--->通过
func numSpecialEquivGroups(A []string) int {
	m := make(map[string]bool)
	for _, a := range A {
		r := [2][]rune{}
		for i, c := range a {
			r[i&1] = append(r[i&1], c)
		}
		sort.Slice(r[0], func(i, j int) bool {
			return r[0][i] < r[0][j]
		})
		sort.Slice(r[1], func(i, j int) bool {
			return r[1][i] < r[1][j]
		})
		m[string(append(r[0], r[1]...))] = true
	}
	return len(m)
}
