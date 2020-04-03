const (
	int32Max = 1 << 31 -1
	int32Min = -1 << 31
)
func myAtoi(str string) int {
	n := len(str)
	flag := 0
	i := 0
	//去掉空格
	for i < n && str[i] == ' '{
		i++
	}
	//标记负号
	if  i < n && str[i] == '-' {
		flag = 1
	}
	//走过正负号
	if i < n && (str[i] == '-' || str[i] == '+') {
		i++
	}
	res := 0
	result := 0
	for i < n && str[i] >= '0' && str[i] <= '9' {
		res = int(str[i] - '0')
		if result > int32Max / 10 || (result == int32Max / 10 && res > 7) {
			if flag == 1 {
				return int32Min
			} else {
				return int32Max
			}
		}
		result = result * 10 + res
		i++
	}
	if flag == 0 {
		return result
	}
	return -result
}