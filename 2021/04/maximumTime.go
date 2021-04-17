package main

//1736. 替换隐藏数字得到的最晚时间

//分清情况即可
func maximumTime(s string) string {
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			//第一位，第二位4-9的情况下第一位只能是1
			if i == 0 && s[i+1] >= '4' && s[i+1] <= '9' {
				res = append(res, byte('1'))
				//第一位，其他情况0-3情况第一位为2
			} else if i == 0 {
				res = append(res, byte('2'))
				//第二位，第一位0或者1的情况下，第二位选择9
			} else if i == 1 && (s[i-1] == '0' || s[i-1] == '1') {
				res = append(res, byte('9'))
				//第二位，第一位2的情况下，第二位选择3
			} else if i == 1 {
				res = append(res, byte('3'))
				//第三位，最大5
			} else if i == 3 {
				res = append(res, byte('5'))
				//第四位，最大9
			} else {
				res = append(res, byte('9'))
			}
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
