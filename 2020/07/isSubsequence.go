package main

import "fmt"

func isSubsequence(s string, t string) bool {
	j := 0
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
		}

		if j == len(s) {
			return true
		}
	}
	return j == len(s)
}

func main(){
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
