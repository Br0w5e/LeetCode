package main


//stack 判断深度
// 当我们遇到一个左括号 ( 时，我们将深度加一，并且新的深度的得分置为 0。当我们遇到一个右括号 ) 时，我们将当前深度的得分乘二并加到上一层的深度。这里有一种例外情况，如果遇到的是 ()，那么只将得分加一。

func scoreOfParentheses(S string) int {
	stack := make([]int, 1)
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*v, 1)
		}
	}
	return stack[len(stack)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
