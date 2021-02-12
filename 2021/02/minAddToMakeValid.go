package main

//921. 使括号有效的最少添加

//栈
func minAddToMakeValid(S string) int {
	res := 0
	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == '('{
			stack = append(stack, S[i])
		} else if len(stack) > 0 && S[i] == ')' && stack[len(stack)-1] == '('{
			stack = stack[:len(stack)-1]
		} else if S[i] == ')' {
			res++
		}
	}
	return res+len(stack)
}

func minAddToMakeValid1(S string) int {
	//left, right 分别表示左右未配对的括号
	left, right := 0, 0
	for _, s := range S {
		if s == '(' {
			left++
		} else if left > 00 && s == ')' {
			left--
		} else {
			right++
		}
	}
	return left+right
}