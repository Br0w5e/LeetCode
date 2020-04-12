//5381. 查询带键的排列
//模拟运行，使用copy函数
func processQueries(queries []int, m int) []int {
    p := make([]int, m)
    for i := 0; i < m; i++{
        p[i]= i+1
    }
    res := make([]int, len(queries))
    for i := 0; i < len(queries); i++{
        index := getIndex(queries[i], p)
        res[i] = index
        tmp := p[index]
        copy(p[1:index+1], p[0:index])
        p[0] = tmp
        
    }
    return res
}

func getIndex(num int, nums []int) int {
    for i, n := range nums {
        if num == n {
            return i
        }
    }
    return -1
}