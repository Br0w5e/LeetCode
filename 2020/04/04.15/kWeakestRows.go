//战斗力最弱的k行
//计算1的个数
//乘100（士兵小于等于100个）加行号
//排序，返回结果
func kWeakestRows(mat [][]int, k int) []int {
    res := make([]int, len(mat))
    for i, row := range mat {
        res[i] = countOne(row)*100 + i
    }
    sort.Ints(res)
    for j := 0; j < k; j++{
        res[j] = res[j]%100
    }
    return res[0:k]

}

func countOne(nums []int) int {
    res := 0
    for _, num := range nums {
        res += num
    }
    //返回1的个数
    return res
}