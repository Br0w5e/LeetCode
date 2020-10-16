package main

//括号的最大嵌入深度
//方法左括号加一，右括号减一，并记录过程中的最大值
//5535. 括号的最大嵌套深度

func maxDepth2(s string) int {
	res, cnt := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		}
		if s[i] == ')' {
			//现有完整括号取最大值
			res = max(res, cnt)
			cnt--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}