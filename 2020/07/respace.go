package main

import "fmt"

//恢复空格
func respace(dictionary []string, sentence string) int {
	dic := make(map[string]int)
	for _, v := range dictionary {
		dic[v] = 1
	}
	n := len(sentence)
	dp := make([]int, n+1) //dp[0]代表句子是空字符串时，没有未识别的字符
	//dp[i]代表前i个字符中最少的未识别字符数
	//dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1      //先假定新加这个字符之后没有匹配到任何单词
		for j := 0; j < i; j++ { //遍历前i个字符，从0到i-1，看从j到i,不包括索引i 组成的子字符串是否能匹配上字典里的单词，如果能匹配，则，dp[j]就是这时的 未匹配字符数，因为从j到i-1已经完全匹配到单词了
			tmp := sentence[j:i]
			if _, ok := dic[tmp]; ok {
				dp[i] = min(dp[i], dp[j])
			}
			if dp[i] == 0 {
				break
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main()  {
	sentence := "jesslookedjustliketimherbrother"
	dictionary := []string{"looked","just","like","her","brother"}
	fmt.Println(respace(dictionary, sentence))
}
