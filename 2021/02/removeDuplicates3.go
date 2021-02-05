package main

//1047. 删除字符串中的所有相邻重复项

func removeDuplicates(S string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		stack = append(stack, S[i])
		for len(stack) > 1 && stack[len(stack)-1] == stack[len(stack)-2] {
			stack = stack[:len(stack)-2]
		}
	}
	return string(stack)
}
