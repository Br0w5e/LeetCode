package main

//字符串相加
func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	l1, l2 := len(num1), len(num2)
	maxLen := max(l1, l2)
	//最长产生maxLen+1 位
	res := make([]byte, maxLen+1)
	for i, j, k := l1-1, l2-1, maxLen; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		var n1, n2 byte
		if i >= 0 {
			n1 = num1[i] - '0'
		}
		if j >= 0 {
			n2 = num2[j] - '0'
		}
		sum := res[k] + n1 + n2
		//本位
		res[k] = sum % 10
		//进位
		res[k-1] = sum / 10
	}
	//最后
	if res[0] == 0 {
		res = res[1:]
	}

	//处理成字符串
	for i := 0; i < len(res); i++ {
		res[i] += '0'
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
