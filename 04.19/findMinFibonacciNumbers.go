//和为K的最小的斐波那契数字数数目
//生成小于等于K的斐波那契数组
//然后进行贪心计算
func findMinFibonacciNumbers(k int) int {
    res := 0
    fibNums := fib(k)
    for i := len(fibNums)-1; i >= 0; i--{
        for k >= fibNums[i] {
            k -= fibNums[i]
            res++
            if k == 0 {
                return res
            }
        }
    }
    return res
}
func fib(n int) []int {
    res := make([]int, 0)
    a, b := 1, 1
    for a <= n{
        res = append(res, a)
        a, b = b, (a+b)
    }
    return res
}