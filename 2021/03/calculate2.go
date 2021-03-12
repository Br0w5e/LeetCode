package main
//227. 基本计算器 II
//面试题 16.26. 计算器

//参考 calculate、evalRPN
//栈实现
func calculate2(s string) int {
	stack := make([]int, 0)
	num := 0
	//非负整数，所以第一个符号为正
	sign := '+'
	for i := 0; i < len(s); i++ {
		if isNum(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		//先算乘除法
		if (!isNum(s[i]) && s[i] != ' ') || i == len(s)-1 {
			switch sign {
			case '-':
				//入栈相反数
				stack = append(stack, -num)
			case '+':
				//入栈
				stack = append(stack, num)
			case '*':
				//出栈
				n1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//直接计算结果放在stack
				stack = append(stack, n1*num)
			case '/':
				//出栈
				n1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//直接计算结果放在stack
				stack = append(stack, n1/num)
			}
			//下一次的符号
			sign = rune(s[i])
			num = 0
		}
	}
	res := 0
	//加法
	for _, e := range stack {
		res += e
	}
	return res
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}
