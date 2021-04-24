package main

//1175. 质数排列

import (
	"fmt"
	"math"
)

func numPrimeArrangements(n int) int {
	all := 0
	for i := 2; i <= n; i++ {
		if judge(i) {
			all++
		}
	}
	return (mul(all) * mul(n-all)) % (int(1e9) + 7)
}

//判断
func judge(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//阶乘
func mul(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
		res %= (int(1e9) + 7)
	}
	return res
}

func main() {
	n := 5
	fmt.Println(numPrimeArrangements(n))
}
