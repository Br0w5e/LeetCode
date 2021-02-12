package main
//1249. 移除无效的括号

//栈和哈希表进行结合
func minRemoveToMakeValid(s string) string {
	//用来匹配括号
	stack := make([]int, 0)
	//记录需要移除的索引
	remove := make(map[int]bool, 0)
	for k, v := range s {
		//左括号入栈
		if v == '(' {
			stack = append(stack, k)
			//右括号，如果栈不为空，则出栈一个，否则该右括号就是多余的，需要标记
		} else if v == ')' {
			//有匹配的左括号，出栈左括号就行
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				//没有匹配的左括号，需要标记移除括号
				remove[k] = true
			}
		}
	}
	//可能左括号存在多余的，将多余的左括号进行标记
	for _, v := range stack {
		remove[v] = true
	}
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		//不是多余的
		if !remove[i] {
			res = append(res, s[i])
		}
	}
	return string(res)
}
