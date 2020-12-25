package main

//455. 分发饼干
import "sort"

//map+模拟--->超时
func findContentChildren(g []int, s []int) int {
	m := make(map[int]int)
	max := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if s[i] > max {
			max = s[i]
		}
	}
	res := 0
	for i := 0; i < len(g); i++ {
		if m[g[i]] > 0 {
			res++
			m[g[i]]--
		} else {
			for j := g[i]; j <= max; j++ {
				if m[j] > 0 {
					res++
					m[j]--
					break
				}
			}
		}
	}
	return res
}
//贪心
func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for i, j := 0, 0; i < len(g) && j < len(s); i++ {
		for j < len(s) && s[j] < g[i] {
			j++
		}
		if j < len(s) {
			res++
			j++
		}
	}
	return res
}
