//求n个不同的数，使得他们和为0
//方法1：对称输出
func sumZero(n int) []int {
    res := make([]int, 0)
    if n%2 == 1{
        for i := 0; i < n; i++{
            res = append(res, i-n/2)
        }
    } else {
        for i := 1; i < n/2+1; i++ {
            res = append(res, i)
            res = append(res, i-n/2-1)
        }
    }
    return res
}
//方法2：前n个正数，最后一个反向求和
func sumZero(n int) []int {
    res := make([]int, 0)
    if n == 0 {
        return nil
    }
    sum := 0 
    for i := 0; i < n-1; i++{
        res = append(res, i)
        sum += i
    }
    res = append(res, -sum)
    return res
}