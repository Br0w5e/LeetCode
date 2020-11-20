package main
//1078. Bigram 分词

//线分割然后再遍历
import "strings"

func findOcurrences(text string, first string, second string) []string {
	s := strings.Split(text, " ")
	if len(s) < 3 {
		return []string{}
	}
	res := make([]string, 0)
	for i := 0; i < len(s)-2; i++ {
		if s[i] == first && s[i+1] == second {
			res = append(res, s[i+2])
		}
	}
	return res
}
