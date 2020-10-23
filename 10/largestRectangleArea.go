package main

//84. 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//方法二
//左边比当前低的第一个
func getLeft(heights []int, idx int) int {
	left := idx
	for i := idx - 1; i >= 0; i-- {
		if heights[i] < heights[idx] {
			break
		}
		left = i
	}
	return left
}
//左边比当前低的第一个
func getRight(heights []int, idx int) int {
	right := idx
	for i := idx + 1; i < len(heights); i++ {
		if heights[i] < heights[idx] {
			break
		}
		right = i
	}
	return right
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		left := getLeft(heights, i)
		right := getRight(heights, i)
		area := (right - left + 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

//遍历
func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		minH := heights[i]
		for j := i; j < len(heights); j++ {
			if minH > heights[j] {
				minH = heights[j]
			}
			if max < minH*(j-i+1) {
				max = minH * (j - i + 1)
			}
		}
	}
	return max
}