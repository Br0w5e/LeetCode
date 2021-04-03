package main

import "fmt"

func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	for _, word := range knowledge {
		m[word[0]] = word[1]
	}

	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		//遇到左括号
		if s[i] == '(' {
			//下一位开始计算，做起始边际
			left := i + 1
			//直到右边界
			for i < len(s) && s[i] != ')' {
				i++
			}
			//不用加一，会在for循环中处理
			word := s[left:i]

			//有的话，使用mao中的val进行替换，否则就用？号进行替换
			if val, ok := m[word]; !ok {
				res = append(res, '?')
			} else {
				res = append(res, []byte(val)...)
			}
			//遇到普通字母，直接进行添加
		} else {
			res = append(res, s[i])
		}
	}
	//返回最终结果
	return string(res)
}

func main() {
	s := "(name)is(age)yearsold"
	knowledge := [][]string{{"name", "bob"}, {"age", "two"}}
	fmt.Println(evaluate(s, knowledge))
}
