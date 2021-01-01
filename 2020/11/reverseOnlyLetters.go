package main

func reverseOnlyLetters(S string) string {
	//最终结果
	res := make([]byte, len(S))
	//记录字符串
	tmp := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		//字符串直接记录
		if (S[i] >= 'A' && S[i] <= 'Z') || (S[i] >= 'a' && S[i] <= 'z') {
			tmp = append(tmp, S[i])
		} else {
			//非字符串进行记录入结果
			res[i] = S[i]
		}
	}
	for i, j := 0, len(tmp)-1; i < len(S) && j >= 0; i++ {
		//将没记录入结果的位置替换成字符串
		if res[i] == 0 {
			res[i] = tmp[j]
			j--
		}
	}
	return string(res)
}