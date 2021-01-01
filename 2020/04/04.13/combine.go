//返回n以内，k个数的所有组合
var result [][]int
func combine(n int, k int) [][]int {
    result = [][]int{}
    if n == 0 || k == 0 {
        return result
    }
    
    nums := []int{}
    process(1, n, k, nums)
    return result
}

func process(start, n, k int, nums []int) {
    if len(nums) == k {
        tmp := make([]int, k)
        copy(tmp, nums)
        result = append(result, tmp)
        return
    }

    for i := start; i <=n-k+len(nums)+1; i++ {
        nums = append(nums, i)
        process(i+1, n, k , nums)
        nums = nums[0:len(nums)-1]
    }

}