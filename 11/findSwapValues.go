package main

//面试题 16.21. 交换和
func findSwapValues(array1 []int, array2 []int) []int {
	sum1 := sum(array1)
	sum2 := sum(array2)
	diff := sum1 - sum2
	//两数之和肯定为偶数 --> 两数之差也为偶数
	if diff % 2 != 0 {
		return []int{}
	}
	m := make(map[int]bool)
	for _, num := range array2 {
		m[num] = true
	}
	for _, num := range array1 {
		if m[num-diff/2] {
			return []int{num, num-diff/2}
		}
	}
	return []int{}
}

func sum(nums []int) int {
	sumNum := 0
	for _, num := range nums {
		sumNum += num
	}
	return sumNum
}