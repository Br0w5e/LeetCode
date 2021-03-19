package main


//面试题 17.19. 消失的两个数字


//map
func missingTwo(nums []int) []int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	res := make([]int, 0)
	for i := 1; i <= len(nums)+2; i++ {
		if m[i] {
			continue
		}
		res = append(res, i)
	}
	return res
}
