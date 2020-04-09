//解码字母到整数的映射
//方法：遍历先考虑3个的，再考虑一个的
func freqAlphabets(s string) string {
	res := make([]byte, 0)
	n := len(s)
	move := 0 
	for i := 0 i < n; {
		if i+2 < n; && s[i+2] == '#' {
			move = int((s[i]-'0'))*10 + int(s[i+1]-'0') - 1
			i += 3
		} else {
			move = int(s[i]-'1')
			i++
		}
		res = append(res, byte(int('a') + move))
	}
	return string(res)
}