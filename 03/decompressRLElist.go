//对压缩信息列表进行解压
func decompressRLElist(nums []int) []int {
    res := make([]int, 0)
    for i := 0; i < len(nums)-1; i+=2{
        for j := 0; j < nums[i]; j++{
            res = append(res, nums[i+1])
        }
    }
    return res
}