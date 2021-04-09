package main

// 26. 删除有序数组中的重复项

//原地删除重复出现的元素
func removeDuplicates(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return n
    }
    slow, fast := 1, 1
    for fast < n {
        if nums[slow-1] != nums[fast] {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}