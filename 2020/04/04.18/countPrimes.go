//计算小于n的素数的个数
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	mark := make([]bool, n)
	for i := 3; i*i < n; i += 2 {
		if !mark[i] {
			// 从 i * i 开始设置，i*2、i*3 等已经在前面被设置过
			for j := i * i; j < n; j += i << 1 {
				mark[j] = true
			}
		}
	}
	total := 1
	for i := 3; i < n; i += 2 {
		if !mark[i] {
			total++
		}
	}
	return total
}

//超时
func countPrimes(n int) int {
    res := 0
    for i := 2; i < n; i++ {
        if isPrime(i) {
            res++
        }
    }
    return res
}

func isPrime(n int) bool {
    for i := 1; i*i <= n; i++{
        if n%i == 0 && i != 1{
            return false
        }
    }
    return true
}