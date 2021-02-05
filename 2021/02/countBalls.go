package main

//5654. 盒子中小球的最大数量

//数组记录即可
func countBalls(lowLimit int, highLimit int) int {
	boxs := make([]int, 60)
	for i := lowLimit; i <= highLimit; i++ {
		sum, tmp := 0, i
		for tmp != 0 {
			sum += tmp%10
			tmp /= 10
		}
		boxs[sum]++
	}
	res := 0
	for _, box := range boxs {
		if box > res {
			res = box
		}
	}
	return res
}
