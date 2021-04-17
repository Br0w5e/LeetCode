package main

//1700. 无法吃午餐的学生数量

//统计学生中1和0的数量，直到栈顶的元素大家都不喜欢
func countStudents(students []int, sandwiches []int) int {
	cnt := make([]int, 2)
	for _, student := range students {
		cnt[student]++
	}
	for _, sandwiche := range sandwiches {
		if cnt[sandwiche] > 0 {
			cnt[sandwiche]--
		} else {
			break
		}
	}
	return cnt[0] + cnt[1]
}
