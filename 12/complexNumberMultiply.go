package main

//537. 复数乘法
import "fmt"

func complexNumberMultiply(a string, b string) string {
	var p, q, m, n int
	fmt.Sscanf(a, "%d+%di", &p, &q)
	fmt.Sscanf(b, "%d+%di", &m, &n)
	return fmt.Sprintf("%d+%di", (p*m-q*n), (m*q+p*n))
}
