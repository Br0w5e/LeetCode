//原地删除重复出现的元素
func removeDuplicates(nums []int) int {
    i := 1
    for j := 2; j < len(nums); j++{
        if nums[j] != nums[i-1]{
            i++
            nums[i]=nums[j]
        }
    }
    return i+1
}