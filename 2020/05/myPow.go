//计算数值的整数次方
//方法一：暴力法-->超时
package main
func myPow(x float64, n int) float64 {
    flag := 0
    if n < 0{
        flag = 1
    }
    if flag == 1{
        return 1/(pow(x, -n))
    }
    return pow(x, n)
}

func pow(x float64, n int) float64{
    res := 1.0
    for i := 0; i < n; i++{
        res *= x
    }
    return res
}

func myPow(x float64, n int) float64 {
	var res = 1.0
	for i := 0; i < abs(n); i++ {
		res *= x
	}
	if n > 0 {
		return res
	}
	return 1/res
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}


//递归
func myPow(x float64, n int) float64 {
    switch n {
        case 0: return 1
        case 1: return x
        case -1: return 1/x
    }
    quotient := myPow(x, n / 2)
    reminder := myPow(x, n % 2)
    return quotient*quotient*reminder
}

//迭代
func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		x = 1 / x
	}
	for n != 0 {
		if n & 1 == 1 {
			res = res * x
		}
		x = x * x
		n = n / 2
	}
	return res
}