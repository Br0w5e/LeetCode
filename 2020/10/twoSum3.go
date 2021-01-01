package main
//1. 两数之和
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if k, ok := m[target-num]; ok {
			return []int{i, k}
		}
		m[num] = i
	}
	return nil
}
