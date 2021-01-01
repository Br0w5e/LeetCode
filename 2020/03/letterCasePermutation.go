//字母大小写全排列
//DFS
func letterCasePermutation(S string) []string {
	res := []string{}
	dfs([]byte(S), 0, &res)
	return res
}
func dfs(S []byte, n int, res *[]string) {
	if n == len(S) {
		*res = append(*res, string(S))
		return
	}
	dfs(S, n+1, res)
	if (S[n] >= 'A' && S[n] <= 'Z') || (S[n] >= 'a' && S[n] <= 'z') {
		//大小写互换
		S[n] ^= 32
		dfs(S, n+1, res)
	}
}