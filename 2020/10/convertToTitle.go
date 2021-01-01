package main

import "fmt"

//168. Excel表列名称
//数字转化为Excel表名

func convertToTitle(n int) string {
	var res string
	for n != 0 {
		//操作很重要
		n--
		res = string(n%26+'A')+res
		n /= 26
	}
	return res
}

func main()  {
	fmt.Println(convertToTitle(27))
}