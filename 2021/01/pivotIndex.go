package main

//724. 寻找数组的中心索引

//前缀和
func pivotIndex(nums []int) int {
	allSum := 0
	for _, num := range nums {
		allSum += num
	}

	halfSum := 0
	for i, num := range nums {
		if halfSum == allSum-halfSum-num {
			return i
		}
		halfSum += num
	}
	return -1
}
