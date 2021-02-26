package main

//5686. 移动所有球到每个盒子所需的最小操作数

//模拟
func minOperations(boxes string) []int {
	res := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			if boxes[j] == '1' {
				res[i] += abs(i-j)
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
