package main

//1539. 第 k 个缺失的正整数

//一次遍历
func findKthPositive(arr []int, k int) int {
	flag := 1
	for i := 0; i < len(arr); i++ {
		for flag < arr[i] {
			flag++
			k--
			if k == 0 {
				return flag-1
			}
		}
		flag++
	}
	return flag-1+k
}
