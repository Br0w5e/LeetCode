package main
//数字转字符串 隔三位加点
//模拟
func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	tmp_string := make([]byte, 0)
	for n != 0 {
		tmp_string = append(tmp_string, byte(n%10+int('0')))
		n /= 10
	}
	left, right := 0, len(tmp_string)-1
	for left < right {
		tmp_string[left], tmp_string[right] = tmp_string[right], tmp_string[left]
		left++
		right--
	}
	res := ""
	for i := len(tmp_string); i >= 0; i = i-3 {
		if i-3 > 0 {
			res = "." + string(tmp_string[i-3:i]) + res
		} else {
			res = string(tmp_string[:i]) + res
		}
	}
	return res
}

