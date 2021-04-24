package main

//1184. 公交站间的距离

//模拟
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	all, halft := 0, 0
	//起始点大于目标点
	if start > destination {
		destination, start = start, destination
	}
	for i, dis := range distance {
		if i >= start && i < destination {
			halft += dis
		}
		all += dis
	}
	if all-halft < halft {
		return all - halft
	}
	return halft
}
