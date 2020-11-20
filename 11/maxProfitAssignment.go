package main

//826. 安排工作以达到最大收益
//贪心
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	//开辟数组，以难度值为长度
	max := 0
	for _, diff := range difficulty {
		if diff > max {
			max = diff
		}
	}
	table := make([]int, max+1)

	for i := 0; i < len(difficulty); i++ {
		diff := difficulty[i]
		if profit[i] > table[diff] {
			table[diff] = profit[i]
		}
	}

	for i := 1; i <= max; i++ {
		if table[i] < table[i-1] {
			table[i] = table[i-1]
		}
	}

	res := 0
	for i := 0; i < len(worker); i++ {
		ability := worker[i]
		if ability > max {
			ability = max
		}
		res += table[ability]
	}
	return res
}
