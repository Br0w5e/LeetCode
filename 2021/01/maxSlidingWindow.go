package main

import "fmt"

//滑动窗口中的最大值

//一次遍历法
func maxSlidingWindow0(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return nil
	}

	left, right := 0, k
	maxIdx := -1
	res := make([]int, 0)
	for i := 0; i < l-k+1; i++ {
		if maxIdx < left {
			maxIdx = left
			for j := left; j < right; j++ {
				if nums[j] > nums[maxIdx] {
					maxIdx = j
				}
			}
		} else if nums[right-1] > nums[maxIdx] {
			maxIdx = right - 1
		}
		res = append(res, nums[maxIdx])
		left++
		right++
	}
	return res
}

//优化后的一次遍历
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	left, right := 0, k
	maxIdx := -1
	res := make([]int, 0)
	for i := 0; i < n-k+1; i++ {
		if maxIdx < left {
			maxIdx = left + MaxIndx(nums[left:right], k)
		} else if nums[right-1] > nums[maxIdx] {
			maxIdx = right - 1
		}
		res = append(res, nums[maxIdx])
		left++
		right++
	}
	return res
}

func MaxIndx(nums []int, k int) int {
	left, right := 0, k-1
	for left < right {
		if nums[left] > nums[right] {
			right--
		} else {
			left++
		}
	}
	return left
}

func maxSlidingWindow2(nums []int, k int) []int {
	window := make([]int, 0)
	res := make([]int, 0)
	for i, num := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= num {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

func main()  {
    nums := []int{1,3,-1,-3,5,3,6,7}
    fmt.Println(maxSlidingWindow2(nums, 3))
}