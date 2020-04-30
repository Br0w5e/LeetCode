//计算海明距离
// 思路：将x和y进行异或结果就是不同的位数，再计算里边1的个数
//优化后算法
func hammingDistance(x int, y int) int {
    n := x^y
    res := 0
    for n != 0 {
        res += n%2
        n >>= 1
    }
    return res
}