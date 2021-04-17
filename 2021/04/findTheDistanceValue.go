package main

//1385. 两个数组间的距离值

//模拟
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, a1 := range arr1 {
		flag := false
		for _, a2 := range arr2 {
			if abs(a1-a2) <= d {
				flag = true
				break
			}
		}
		if !flag {
			res++
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
