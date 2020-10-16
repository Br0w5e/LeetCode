package main

//1608. 特殊数组的特征值

func specialArray(nums []int) int {
	//i用来遍历所有可能结果
	for i := 0; i < len(nums)+1; i++ {
		tmp := 0
		//统计数组
		for _, num := range nums {
			if num >= i {
				tmp++
			}
		}
		if tmp == i {
			return i
		}
	}
	return -1
	
}
