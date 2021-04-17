package main

//678. 有效的括号字符串

func checkValidString(s string) bool {
	//左括号的位置
	left := make([]int, 0)
	//星号的位置
	star := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left = append(left, i)
		} else if s[i] == '*' {
			star = append(star, i)
			//遇到右括号
		} else {
			//没有与之匹配的左括号或者星
			if len(left) == 0 && len(star) == 0 {
				return false
			}
			//优先考虑左括号
			if len(left) > 0 {
				left = left[:len(left)-1]
				//接着考虑星
			} else {
				star = star[:len(star)-1]
			}
		}
	}
	//左括号数量一定得为0， 不允许星号出现在左括号之后
	for len(left) > 0 && len(star) > 0 {
		if left[len(left)-1] > star[len(star)-1] {
			return false
		}
		left = left[:len(left)-1]
		star = star[:len(star)-1]
	}
	//最后左括号一定要为0
	return len(left) == 0
}
