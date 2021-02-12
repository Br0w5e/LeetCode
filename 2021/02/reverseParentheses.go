package main

//1190. 反转每对括号间的子串

//将左括号的位置入栈，每次遇到右括号翻转相应范围，并将对应左括号出栈
func reverseParentheses(s string) string {
	stack := make([]int, 0)
	bytes := []byte(s)
	for k, v := range s {
		if v == '(' {
			stack = append(stack, k)
		}
		if v == ')' {
			left := stack[len(stack)-1]
			reverseBytes(bytes, left, k)
			stack = stack[:len(stack)-1]
		}
	}
	res := make([]byte, 0)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ')' || bytes[i] == '(' {
			continue
		}
		res = append(res, bytes[i])
	}
	return string(res)
}

func reverseBytes(bytes []byte, left int, right int) {
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}
