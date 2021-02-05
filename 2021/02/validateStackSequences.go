package main

//946. 验证栈序列

//设置临时栈
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) >= 1 && len(popped) >= 1 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(popped) == 0
}
