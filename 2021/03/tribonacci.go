package main

//1137. 第 N 个泰波那契数

//仿照斐波那契，迭代
func tribonacci(n int) int {
	T0, T1, T2 := 0, 1, 1
	if n == 0 {
		return T0
	} else if n== 1 {
		return T1
	}
	for i := 3; i <= n; i++ {
		T2, T1, T0 = T2+T1+T0, T2, T1
	}
	return T2
}