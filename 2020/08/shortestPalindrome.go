package main

import "fmt"

func shortestPalindrome(s string) string {
	rs := make([]byte, 0)
	l := len(s)
	//逆序
	for i := l-1; i >=0; i-- {
		rs = append(rs, s[i])
	}
	rss := string(rs)
	flag := 0
	//求最长后缀
	for rss[flag:] != s[:l-flag]{
		flag++
	}
	//最短前缀相加
	return rss[:flag]+s
}

func main()  {
	s := "aacecaaa"
	fmt.Println(shortestPalindrome(s))
}