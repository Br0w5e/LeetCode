package main

//204. 计数质数
import "fmt"

//计算小于n的素数的个数
//筛选法
func countPrimes(n int) int {
	m := make([]bool, n)
	for i := 2; i*i <= n; i++ {
		if !m[i] {
			for j := i * i; j < n; j += i {
				m[j] = true
			}
		}
	}

	res := 0
	for i := 2; i < n; i++ {
		if !m[i] {
			res++
		}
	}
	return res
}

//超时
func countPrimes2(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func isPrime(n int) bool {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 && i != 1 {
			return false
		}
	}
	return true
}

func main()  {
	fmt.Println(countPrimes(100))
}