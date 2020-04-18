//判断给定数字是不是2的幂
//求解
func isPowerOfTwo(n int) bool {
    res := 1
    for res < n {
        res *= 2
    }
    return res == n
}

//位运算
func isPowerOfTwo(n int) bool {
    if n == 0 {
        return false
    }
    return n & (n-1) == 0
}