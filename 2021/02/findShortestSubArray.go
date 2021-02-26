package main

//697. 数组的度

//哈希表记录出现的小标
func findShortestSubArray(nums []int) int {
	m := make(map[int][]int)
	degree := 1
	res := 1
	for i, v := range nums {
		m[v] = append(m[v], i)
		//当前的度大于，那无论如何都是该数字
		if len(m[v]) > degree {
			degree = len(m[v])
			res = m[v][degree-1] - m[v][0] + 1
			//度相等，判断长度是否小于
		} else if degree == len(m[v]) {
			if m[v][degree-1]-m[v][0]+1 < res {
				res = m[v][degree-1] - m[v][0] + 1
			}
		}
	}
	return res
}
