//计算能接多少雨水
//思路：将柱子填充成左边不严格递增，右边不严格递增的图形，最终增量即为接到的水的数量
//双指针
func trap(height []int) int  {
	n := len(height)
	if n < 3 {
		return 0
	}
	left, right := 0, n-1
	result := 0
	l_max, r_max := height[left], height[right]
	for left < right {
		if l_max < height[left] {
			l_max = height[left]
		}
		if r_max < height[right] {
			r_max = height[right]
		}
		if l_max < r_max {
			result += (l_max - height[left])
			left++
		} else {
			result += (r_max - height[right])
			right--
		}
	}
	return result
}

//暴力
func trap(height []int) int {
    n := len(height)
    left := make([]int, n)
    right := make([]int, n)
    for i := 1; i < n; i++{
        left[i] = max(left[i-1], height[i-1])
    }
    for i := n-2; i >= 0; i--{
        right[i] = max(right[i+1], height[i+1])
    }
    water := 0
    for i := 0; i < n; i++{
        level := min(left[i], right[i])
        water += max(0, level - height[i])
    }
    return water
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}