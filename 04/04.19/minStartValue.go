//逐步求和得到正数的最小值
//写一个judge函数，用来判断该数字能否满足要求。
func minStartValue(nums []int) int {
    res := 1
    for res >= 1 {
        if judge(nums, res) {
            return res
        } else {
            res++
        }
    }
    return 0
}

func judge(nums []int, n int) bool {
    for i := 0; i < len(nums); i++{
        n += nums[i]
        if n < 1 {
            return false
        }
    }
    return true
}