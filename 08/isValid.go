package main

import "fmt"

func isValid(s string) bool {
	c := make([]byte, 0)
	symbols := map[byte]byte {
		')' : '(',
		']' : '[',
		'}' : '{',
	}
	for _, value := range s {
		n := len(c)
		if n > 0 {
			//判断新括号能否与上一个抵消
			if c[n-1] == symbols[byte(value)] {
				c = c[:n-1]
				continue
			}
		}
		c = append(c, byte(value))
	}
	return len(c) == 0
}

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}