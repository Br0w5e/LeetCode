package main

//659. 分割数组为连续子序列

func isPossible(nums []int) bool {
	//统计各个元素出现的次数
	m := make(map[int]int)
	//长度至少为3的结尾情况
	n := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for _, num := range nums {
		if m[num] == 0 {
			continue
		}
		m[num]--

		if n[num-1] > 0 {
			n[num-1]--
			n[num]++
		} else if m[num+1] > 0 && m[num+2] > 0 {
			m[num+1]--
			m[num+2]--
			n[num+2]++
		} else {
			return false
		}
	}
	return true
}
