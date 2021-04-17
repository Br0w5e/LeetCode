package main

//1460. 通过翻转子数组使两个数组相等

//map，所有出现的频率相同就可
func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int)
	for i, v := range target {
		m[v]++
		m[arr[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

