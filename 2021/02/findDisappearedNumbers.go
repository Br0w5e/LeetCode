package main

//448. 找到所有数组中消失的数字

//自身map
func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	n := len(nums)
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

//普通map
func findDisappearedNumbers1(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	for _, v := range nums {
		v = (v-1)%n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}
	return res
}
