package main

//1455. 检查单词是否为句中其他单词的前缀

//先拆分句子，然后计算
import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	sentenceArr := strings.Split(sentence, " ")
	for i, word := range sentenceArr {
		if len(word) >= len(searchWord) && word[:len(searchWord)] == searchWord {
			return i+1
		}
	}
	return -1
}

func main()  {
	sentence := "i love eating burger"


}
