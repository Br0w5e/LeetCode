//生成括号
//深度优先遍历穷举所有可能性，并在过程中进行条件判断（剪枝）
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gen("", 0, 0, n, &res)
	return res
}

func gen(str string, left int, right int, n int, res *[]string) {
	if left == n && right == n {
		*res = append(*res, str)
		return 
	}
	if left < n {
		gen(str+"(", left+1, right, n, res)
	}
	//带剪枝
	if right < left  && right < n {
		gen(str+")", left, right+1, n, res)
	}
}
