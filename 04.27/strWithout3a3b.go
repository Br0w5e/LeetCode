package main

import "fmt"

func strWithout3a3b(A int, B int) string {
	l := A + B
	ans := ""
	for A > 0 || B > 0 {
		if A > B {
			ans += "aab"
			A -= 2
			B -= 1
		} else if A < B {
			ans += "bba"
			A -= 1
			B -= 2
		} else {
			ans += "ab"
			A -= 1
			B -= 1
		}
	}
	return ans[:l]
}

func main()  {
	A, B := 4, 5
	res := strWithout3a3b(A, B)
	fmt.Println(res)
}