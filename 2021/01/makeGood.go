package main
//1544. 整理字符串

//遍历字符串，每次都将遍历放进栈中，并检查栈顶两个元素是否为同一字符的大小写
func makeGood(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		res = append(res, s[i])
		size := len(res)
		if size >= 2 {
			// 判断
			if judge(res[size-1], res[size-2]) {
				res = res[:size-2]
			}
		}
	}
	return string(res)
}


//判断是否为同一字符的大小写
func judge(a, b byte) bool {
	//同一字母大小写异或为32 即为’ ‘
	return a^b == ' '
}
