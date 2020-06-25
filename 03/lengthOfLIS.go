//寻找最长上升子序列
//方法一：动态规划

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++{
		res[i] = 1
	}
	for i := 0; i < len(nums); i++{
		for j := 0; j < i; j++{
			if nums[j] < nums[i] {
				res[i] = max(res[i], res[j]+1)
			}
		}
	}
	return getMax(res)
	
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMax(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return nums[left]
}

//动态规划+二分查找
func lengthOfLIS(nums []int) int {
    tail := make([]int, len(nums))
    res := 0
    for _, num := range nums {
        left, right := 0, res
        for left < right {
            mid := (left+right)/2
            if tail[mid] < num {
                left = mid+1
            } else {
                right = mid
            }
        }
        tail[left] = num
        if right == res {
            res++
        }
    }
    return res
}