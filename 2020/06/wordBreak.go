package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//一种优化后的代码，但是看起来有点复杂
/*
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
    //建表
	for _, word := range wordDict {
		wordDictSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
    //获取最长单词
    maxWordLen := maxLen(wordDict)

	for i := 1; i <= len(s); i++ {
		for j := max(0, i-maxWordLen); j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func maxLen(wordList []string) int {
    max := 0
    for _, word := range wordList {
        if len(word) > max {
            max = len(word)
        }
    }
    return max
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
 */


func main() {
	s := "catscat"
	wordDict :=  []string{"cat", "cats"}
	fmt.Println(wordBreak(s, wordDict))

}