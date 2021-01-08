package main
//509. 斐波那契数

func fib(N int) int {
    if N == 0 {
        return 0
    }
    a, b := 0, 1
    for i := 0; i < N - 1; i++ {
        a, b = b, a+b
    }
    return b
    
}

func fib(n int) int {
    var MOD int = 1e9+7
    a, b := 0, 1
    for n > 0{
        a, b = b%MOD, (a+b)%MOD
        n--
    }
    return a
}

/* 递归的解决方案
var fibs = make([]int, 31)
func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}

	res := fibs[N]
	if res == 0 {
		res = fib(N-1) + fib (N-2)
		fibs[N] = res
	}

	return res
}
*/

/* 剑指offer中的题解
var fibs = make([]int, 101)
func fib(n int) int {
    if n == 0 || n == 1 {
        return n
    }
    res := fibs[n]
    if res == 0{
        res = (fib(n-1) + fib(n-2))%1000000007
        fibs[n] = res
    }

    return res
    
}
*/

//青蛙跳台阶，挖掘出斐波那契这一条件
/*
func numWays(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    a, b := 0, 1
    m := 1000000007
    for i := 0; i < n; i++{
        a, b = b%m, (a+b)%m
    } 

    return b
}
*/