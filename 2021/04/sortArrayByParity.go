package main

//905. 按奇偶排序数组
//双指针
func sortArrayByParity(A []int) []int {
	res := make([]int, len(A))
	left, right := 0, len(res)-1
	for _, a := range A {
		if a % 2 == 0 {
			res[left] = a
			left++
		} else {
			res[right] = a
			right--
		}
	}
	return res
}
