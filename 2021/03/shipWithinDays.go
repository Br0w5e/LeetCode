package main

//1011. 在 D 天内送达包裹的能力

//二分查找，试结果
func shipWithinDays(weights []int, D int) int {
	min, max := -1, 0
	for _, weight := range weights {
		max += weight
		if weight > min {
			min = weight
		}
	}

	for min < max {
		mid := min + (max-min) >> 1
		tmp := 0
		day := 1

		for _, weight := range weights {
			tmp += weight
			if tmp > mid {
				day++
				tmp = weight
			}
		}

		//重点关注这一块，是根据天数来确定质量的，当天数大于给定天数，说明选择质量过小，当天数小于给定质量的时候说明选定质量过大
		if day > D {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}
