package main

//1003. 检查替换后的词是否有效

//遇到abc就出栈，判断最后栈是不是空的
func isValid(s string) bool {
	if len(s) < 3 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		for len(stack) >= 3 && string(stack[len(stack)-3:len(stack)]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}