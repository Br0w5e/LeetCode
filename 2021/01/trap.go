package main
//面试题 17.21. 直方图的水量
//42. 接雨水

//计算能接多少雨水
//思路：将柱子填充成左边不严格递增，右边不严格递减的图形，最终增量即为接到的水的数量
//朴素做法
func trap(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    res := 0
    //寻找最高点
    heightIndex := getMaxIndex(height)
    max := 0

    //最高点以左，非严格递增，增量即为水
    for i := 0; i < heightIndex; i++ {
        if height[i] > max {
            max = height[i]
        }
        res += max-height[i]
    }
    //最高点以右，非严格递减，差值为水
    max = 0
    for i := n-1; i > heightIndex; i-- {
        if height[i] > max {
            max = height[i]
        }
        res += max-height[i]
    }
    return res
}

func getMaxIndex(nums []int) int {
    res := 0
    for _, num := range nums {
        if num > res {
            res = num
        }
    }
    for i, num := range nums {
        if num == res {
            return i
        }
    }
    return -1
}

//双指针
func trap2(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	res := 0
	l, r := 0, n-1
	lmax, rmax := height[l], height[r]
	for l < r {
		if lmax < height[l] {
			lmax = height[l]
		}
		if rmax < height[r] {
			rmax = height[r]
		}
        //尚未到达最高峰，所以左边不严格递增
		if lmax < rmax {
			res += (lmax - height[l])
			l++
		} else {
		    //右边不严格递减
			res += (rmax - height[r])
			r--
		}
	}
	return res
}

//暴力
func trap3(height []int) int {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i+1])
	}
	water := 0
	for i := 0; i < n; i++ {
		level := min(left[i], right[i])
		water += max(0, level-height[i])
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
