//在指定的递增数列中寻找指定的和
//方法：双指针查找法
package  main

import "fmt"

func twoSum(nums []int, target int) []int {
    size := len(nums)
    if len(nums) == 0 {
        return nil
    }
    l, h := 0, size-1
    for l < h {
        if nums[l] + nums[h] < target {
            l++
        } else if nums[l] + nums[h] > target {
            h--
        } else {
            return []int{nums[l], nums[h]}
        }
    } 
    return nil
}
//返回下标
func twoSum1(numbers []int, target int) []int {
    left, right := 0, len(numbers)
    for left < right {
        if numbers[left]+numbers[right] < target {
            left++
        } else if numbers[left]+numbers[right] == target {
            return []int{left+1, right+1}
        } else {
            right--
        }
    }
    return nil
}
