package main
//150. 逆波兰表达式求值
import "strconv"

//栈逐步遍历，遇到符号处理栈里边的前两个
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		//遇到符号
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			res := 0
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			switch {
			case tokens[i] == "+":
				res = n1+n2
			case tokens[i] == "-":
				res = n1-n2
			case tokens[i] == "/":
				res = n1/n2
			case tokens[i] == "*":
				res = n1*n2
			}
			stack = append(stack, res)
		} else {
			//数字入栈
			tokenInt, _ := strconv.Atoi(tokens[i])
			stack = append(stack, tokenInt)
		}
	}
	return stack[len(stack)-1]
}