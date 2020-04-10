//生成集合的幂集
//方法：迭代
//遇到一个数就把所有子集加上该数组成新的子集，遍历完毕即是所有子集。
func subsets(nums []int) [][]int {
    //初始时只有空集
    res := make([][]int, 1)
    //获取每个元素
    for _, num := range nums {
        //拿当前元素与前面所有元素进行组合
        for _, ele := range res {
            tmp := make([]int, len(ele), len(ele)+1)
            copy(tmp, ele)
            tmp = append(tmp, num)
            res = append(res, tmp)
        }
    }
    return res
}