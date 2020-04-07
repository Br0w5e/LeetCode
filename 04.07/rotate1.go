//旋转数组
//将数组顺时针旋转k次
//朴素方法：模拟
func rotate(nums []int, k int) []int {
    n := len(nums)
    k %= n
    if k == 0 {
        return nums
    }
    for i := 0; i < k; i++{
        tmp := nums[n-1]
        for j := n-1; j > 0; j--{
            nums[j] = nums[j-1]
        }
        nums[0] = tmp
    }
    return nums
}

//三次翻转：
func reverse(nums []int) {  // 翻转数组
    left, right := 0, len(nums)-1
    for left < right {
        nums[left],nums[right] = nums[right],nums[left]
        left++
        right--
    }
}

func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    //前n-k翻转
    reverse(nums[:n-k])
    //后n-k翻转
    reverse(nums[n-k:])
    //集体翻转
    reverse(nums)
}