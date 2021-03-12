package main

import "strings"

//1577. 数的平方等于两数乘积的方法数

//暴力---> 超时
func numTriplets(nums1 []int, nums2 []int) int {
	res := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			for k := j+1; k < len(nums2); k++ {
				if nums1[i]*nums1[i] == nums2[j]*nums2[k] {
					res++
				}
			}
		}
	}

	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {
			for k := j+1; k < len(nums1); k++ {
				if nums2[i]*nums2[i] == nums1[j]*nums1[k] {
					res++
				}
			}
		}
	}
	return res
}

//使用map记录
func numTriplets(nums1 []int, nums2 []int) int {
	res := 0
	res += getTripleNum(nums1, nums2)
	res += getTripleNum(nums2, nums1)
	return res
}

func getTripleNum(nums1 []int, nums2 []int) int {
	res := 0
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num*num]++
	}
	for i := 0; i < len(nums2); i++ {
		for j := i+1; j < len(nums2); j++ {
			res += m[nums2[i]*nums2[j]]
		}
	}
	strings.Split()
	return res
}

