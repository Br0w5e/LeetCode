package main

//1491. 去掉最低工资和最高工资后的工资平均值

//模拟
func average(salary []int) float64 {
	min := 1000000
	max := 0
	res := 0
	for _, s := range salary {
		res += s
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}

	res -= (max+min)
	return float64(res)/float64(len(salary)-2)
}