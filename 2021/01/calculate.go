package main

//224. 基本计算器
func calculate(s string) int {
	stack := make([]int, 0)
	flag, res := 1, 0
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '+':
			flag = 1
		case s[i] == '-' :
			flag = -1
		case s[i] == '(':
			//将以前的计算结果添加到栈里边
			stack = append(stack, res)
			res = 0
			//添加上一位正负号
			stack = append(stack, flag)
			flag = 1
		case s[i] == ')':
			//计算括号里的结果
			res = stack[len(stack)-1]*res + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		case s[i] == ' ':
		default:
			num := int(s[i]-'0')
			//出现连续多个数字的情况
			for i+1 < len(s) && int(s[i+1]-'0') >= 0 && int(s[i+1]-'0') <= 9 {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			res += flag*num
		}
	}
	return res
}
