package main

//1130. 叶值的最小代价生成树


//单调递减栈，因为越小的元素深度越大，才能保证使用的次数越多，最后的值也就越大。
func mctFromLeafValues(arr []int) int {
	//单调递减栈
	stack := make([]int, 0)
	//哨兵
	stack = append(stack, 16)
	res := 0
	for i := 0; i < len(arr); i++ {
		//栈顶元素是当前的极小值，arr[i]是当前元素
		for stack[len(stack)-1] < arr[i] {
			//取左右两边较小的乘以栈顶元素
			res += stack[len(stack)-1]*min(arr[i], stack[len(stack)-2])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	//剩余元素
	for len(stack) > 2 {
		res += stack[len(stack)-1]*stack[len(stack)-2]
		stack = stack[:len(stack)-1]
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}