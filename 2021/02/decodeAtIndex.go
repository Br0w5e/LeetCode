package main
//880. 索引处的解码字符串

//栈思想
//先统计字符串长度，然后从后往前遍历，当K对当前长度取余为0，且当前为字母时则为K位字符
func decodeAtIndex(S string, K int) string {
	//字符串总长度
	size := 0
	for i := 0; i < len(S); i++ {
		if isdigit(S[i]) {
			size *= int(S[i]-'0')
		} else {
			size++
		}
	}

	for i := len(S)-1; i >= 0; i-- {
		if isdigit(S[i]) {
			size /= int(S[i]-'0')
			K %= size
		} else {
			if K % size == 0 {
				return string(S[i])
			} else {
				size--
			}
		}
	}
	return ""
}

func isalpha(s byte) bool {
	return s >= 'a' && s <= 'z'
}

func isdigit(s byte) bool {
	return s >= '0' && s <= '9'
}


