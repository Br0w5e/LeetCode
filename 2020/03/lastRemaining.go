//最后一个存活的的人
//方法：约瑟夫环问题
// 从后往前推，最后一个存活的人肯定是0
func lastRemaining(n int, m int) int {
    flag := 0
    for i := 2; i <= n; i++{
        flag = (flag + m) % i
    }
    return flag
}