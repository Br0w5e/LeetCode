package main

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	for i := 0; i < len(T)-1; i++ {
		for j := i+1; j < len(T); j++ {
			if T[j] > T[i] {
				res[i] = j-i
				break
			}
		}
	}
	return res
}

func dailyTemperatures(T []int) []int {

	R := make([]int, len(T))
	stack := []int{}

	for k, v := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < v {
			// i为栈底温度的日期
			i := stack[len(stack)-1]
			//pop
			stack = stack[:len(stack)-1]
			//k-i = 当前温度日期 - 栈底温度日期 = 升高温度日期
			R[i] = k - i
		}
		//入栈的是下标，不是温度，下标对应当前温度日期
		stack = append(stack, k)
	}
	return R
}