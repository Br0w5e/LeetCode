package main

//1052. 爱生气的书店老板

//滑动窗口
func maxSatisfied(customers []int, grumpy []int, X int) int {
	//高兴时刻
	atLeast := 0
	for i, v := range grumpy {
		atLeast += (1-v) * customers[i]
	}

	//改变态度时刻的基础窗口
	change := 0
	for i := 0; i < X; i++ {
		change += customers[i]*grumpy[i]
	}

	//让窗口动起来
	maxChange := change
	for i := X; i < len(grumpy); i++ {
		change += customers[i]*grumpy[i] - customers[i-X]*grumpy[i-X]
		if change > maxChange {
			maxChange = change
		}
	}

	//结果
	return atLeast+maxChange
}
