//计算数字阶乘有多少0
func trailingZeroes(n int) int {    // 算一下因子里有多少个5就可以了,2有很多个就不用管了
    cnt := 0
    for n >= 5 {
        cnt += n/5
        n /= 5
    }
    return cnt
}