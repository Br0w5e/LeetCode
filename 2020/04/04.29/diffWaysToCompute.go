package main

import "strconv"

//给表达式加括号
func diffWaysToCompute(input string) (res []int) {
	for i, v := range input {
		c := string(v)
		if c == "+" || c == "-" || c == "*" {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, i := range left {
				for _, j := range right {
					if c == "+" {
						res = append(res, i+j)
					} else if c == "-" {
						res = append(res, i-j)
					} else {
						res = append(res, i*j)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		num, _ := strconv.Atoi(input)
		res = append(res, num)
	}
	return
}