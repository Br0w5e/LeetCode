package main
//5629. 重新格式化电话号码

//预处理+模拟
func reformatNumber(number string) string {
	tmp := ""
	for i := 0; i < len(number); i++ {
		if number[i] != ' ' && number[i] != '-' {
			tmp += string(number[i])
		}
	}
	res := ""
	i := 0
	for i = 0; i < len(tmp)-4; i+=3 {
		res += tmp[i:i+3]+"-"
	}
	if i == len(tmp)-4 {
		res += tmp[i:i+2]+"-"+tmp[i+2:]
	} else {
		res += tmp[i:]
	}
	return res
}
