package main

//349. 两个数组的交集
//求两个集合的交集

//双哈希表
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	res := make([]int, 0)

	for _, num := range nums1 {
		m1[num] = true
	}
	for _, num := range nums2 {
		m2[num] = true
	}

	for k, _ := range m1 {
		if m2[k] {
			res = append(res, k)
		}
	}
	return res
}

//单哈希表
//注意过程中改变表的值
func intersection2(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)

	for _, num := range nums1 {
		m[num] = true
	}
	for _, num := range nums2 {
		if m[num] {
			res = append(res, num)
			m[num] = false
		}
	}

	return res
}
