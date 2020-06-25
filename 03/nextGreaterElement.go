func nextGreaterElement(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    for _, num := range nums1{
        res = append(res, getFirstBig(nums2, num))
    }
    return res
}

func getFirstBig(nums []int, target int) int {
    n := len(nums)
    i := 0
    for i = 0 ; i < n && nums[i] != target; i++{

    }
    for j := i; j < n;{
        if nums[j] > target{
            return nums[j]
        } else{
            j++
        }
    }
    return -1
}