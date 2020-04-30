//移除最外边的括号
//用flag做标记，当flag==0时候加入
func removeOuterParentheses(S string) string {
	var res string
	flag := 0
	left := 0
    for i, s := range S {
        if s == '(' {
            flag++
        } else {
            flag--
            if flag == 0 {
                res += S[left+1 : i]
                left = i + 1
            }
        }
    }
	return res
}