package main

import "sort"

//哈希表
func intersect(nums1 []int, nums2 []int) []int {
	//if len(nums1) < len(nums2) {
	//	return intersect(nums2, nums1)
	//}
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	res := make([]int, 0)
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}

//排序

func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0
	res := make([]int, 0)
	for index1 < length1 && index2 < length2{
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			res = append(res, nums1[index1])
			index1++
			index2++
		}
	}
	return res
}
