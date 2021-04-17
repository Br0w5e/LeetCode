package main

//1716. 计算力扣银行的钱
func totalMoney(n int) int {
	res := 0
	week := n/7
	day := n%7
	//所有周，完整周
	for i := 0; i < week; i++ {
		//每周
		for j := 0; j < 7; j++ {
			res += j+1+ i
		}
	}
	// 剩余的天
	for i := 0; i < day; i++ {
		res += i+1 + week
	}

	return res
}
