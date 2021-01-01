//对字符串进行压缩
//双指针
func compressString(S string) string {
	if len(S) < 2 {
		return S
	}
	res := ""
	slow, quick := 0, 0
	for quick < len(S) {
		if S[quick] != S[slow] {
			res += string(S[slow]) + strconv.Itoa(quick-slow)
			slow = quick
		}
		quick++
	}
	//最后一个
	res += string(S[slow]) + strconv.Itoa(quick-slow)

	if len(res) >= len(S) {
		return S
	}

	return res
}