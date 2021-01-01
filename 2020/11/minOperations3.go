package main

//5602. 将 x 减到 0 的最小操作数

//滑动窗口
//反向思考一下
import "math"

//数组的和为sum，那么窗口的最大值为win = sum-x。
//
//如果当前窗口w > win时，就尝试即移动窗口的左边沿，缩小窗口，即start增大
//如果当前窗口w == win时，start + nums.length - 1 - i（也就是窗口外的区域大小），就是可能的解

func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum == x {
		return len(nums)
	}
	//剩余窗口的大小
	win := sum - x
	//当前和的大小
	w := 0
	start := 0
	res := math.MaxInt64

	for i := 0; i < len(nums); i++ {
		w += nums[i]
		//当前和过大，尝试缩小窗口
		if w > win {
			for start < i && w > win {
				w -= nums[start]
				start++
			}
		}
		//w大小就是可能值
		if w == win {
			res = min(res, start+len(nums)-1-i)
		}
	}
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
