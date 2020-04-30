package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]int)
	for {
		n := getNum(n)
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		} else {
			m[n]++
		}
	}
}
func getNum(n int) int {
	res := 0
	for n != 0 {
		res += (n%10)*(n%10)
		n /= 10
	}
	return res
}

func main()  {
	n := 10
	for n != 0{
		if isHappy(n) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
		n--
	}
}