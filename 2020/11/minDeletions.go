package main

import "sort"
//5562. 字符频次唯一的最小删除次数
//贪心
func minDeletions(s string) int {
	nums := make([]int, 26)
	for _, ch := range s {
		nums[int(ch-'a')]++
	}
	sort.Ints(nums)
	res := 0
	i := 0
	for i = 24; i >= 0; {
		//提前减少为0，则停止，并统计前面出现大于0的个数之和
		if nums[i] == 0 {
			break
		}
		//操作并后移
		if nums[i] == nums[i+1] {
			nums[i]--
			res++
			i--
			//操作但不后移
		}else if nums[i] > nums[i+1] {
			nums[i]--
			res++
			//不进行操作，但后移
		}else {
			i--
		}
	}
	for j := 0; j < i; j++ {
		res += nums[j]
	}
	return res

}
