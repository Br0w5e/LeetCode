//摆动排序
//开辟tmp数组，并将tmp排序
func wiggleSort(nums []int)  {
    n := len(nums)
    tmp := make([]int, n)
    copy(tmp, nums)
    sort.Ints(tmp)
    j, k := n, (n+1)/2
    for i := 0; i < n; i++{
        //奇数位,大到小
        if i%2 == 1 {
            j--
            nums[i] = tmp[j]
        //偶数位
        } else {
            k--
            nums[i] = tmp[k]
        }
    }
}