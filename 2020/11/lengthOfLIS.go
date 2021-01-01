//寻找最长上升子序列
//方法一：动态规划
package main

//300. 最长上升子序列
//对比673
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
func lengthOfLIS2(nums []int) int {
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

func lengthOfLIS3(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	//dp为到达第i个位置时的最长子序列长度
	dp := make([]int, len(nums))
	//初始值
	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//如果dp[i] < dp[j]+1 表明在0-i-1中，出现了
				//新的最大值，则将最大值更新
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}