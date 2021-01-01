package main
//1573. 分割字符串的方案数
func numWays(s string) int {
	mod := 1e9+7
	n := len(s)
	//记录1的位置
	pos := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			pos = append(pos, i)
		}
	}
	//1的个数不是3的倍数
	if len(pos)%3 != 0 {
		return 0
	}
	//全0 总共有(Cn-1,2)个分割点
	if len(pos) == 0 {
		return ((n-1)*(n-2)/2)%int(mod)
	}
	//是3的倍数,每块里边1的数量是一定的都为sile个
	sile := len(pos)/3
	//pos[sile]-pos[sile-1]：s1字符串和s2字符串中可分割的位置
	//pos[sile*2]-pos[sile*2-1]：s2字符串和s3字符串中可分割的位置
	return (pos[sile]-pos[sile-1])*(pos[sile*2]-pos[sile*2-1])%int(mod)
}
