package main

import (
	"fmt"
	"math"
)

//507. 完美数
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	res := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			res += i
			res += num/i
		}
	}
	return res == num
}
func main()  {
	num := 6
	fmt.Println(checkPerfectNumber(num))
}