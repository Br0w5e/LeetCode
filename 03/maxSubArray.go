//连续最大的子序列
//动态规划算法
func maxSubArray(nums []int) int {
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > 0 {
            nums[i] += nums[i-1]
        }
        //好好理解这
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}