package main

//779. 第K个语法符号
//暴力模拟---> 超时
func kthGrammar(N int, K int) int {
	pre, cur := "0", ""
	for i := 1; i <= N; i++ {
		for j := 0; j < len(pre); j++ {
			if pre[j] == '0' {
				cur += "01"
			} else {
				cur += "10"
			}
		}
		pre, cur = cur, ""
	}
	return int(pre[K-1]-'0')
}

//递归

//第 K 个数字是上一行第 (K+1) / 2 个数字生成的。如果上一行的数字为 0，被生成的数字为 1 - (K%2)，如果上一行的数字为 1，被生成的数字为 K%2
func kthGrammar(N int, K int) int {
	if N == 1 {
		return (K+1)%2
	}
	return kthGrammar(N-1, (K+1)/2)^((K+1)%2)
}

