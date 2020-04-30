//重复N次的数字
//方法：排序-->相邻的
func repeatedNTimes(A []int) int {
    sort.Ints(A)
    for i := 0; i < len(A); i++{
        if A[i] == A[i+1]{
            return A[i]
        }
    }
    return -1
}
//map
func repeatedNTimes(A []int) int {
    m := make(map[int]int)
    for _, num := range A {
        m[num]++
    }
    for k, v := range m {
        if v != 1 {
            return k
        }
    }
    return -1
}
//不排序直接算
func repeatedNTimes(A []int) int {
    //隔一个
    if A[1] == A[3] {
        return A[1]
    }
    //相邻的
    for i := 0; i< len(A)-1; i++{
        if A[i] == A[i+1] {
            return A[i]
        }
    }
    //第一个
    return A[0]
}