//求可装水的最大容量
//方法一：两次遍历，并将每一轮的结果存储在数组中，最后寻求数组的最大值。
func maxArea(height []int) int {
    n := len(height)
    res := make([]int, n-1)
    for i := 0; i < n-1; i++{
        maxNum := 0
        for j := i+1; j < n; j++{
            if min(height[i], height[j]) * (j-i) > maxNum {
                maxNum = min(height[i], height[j]) * (j-i)
            }
        }
        res[i] = maxNum
    }
    return getMax(res)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
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
//方法二：沿用方法一，但不开辟数组，每次调整最大值
func maxArea(height []int) int {
    n := len(height)
    maxNum := 0
    for i := 0; i < n-1; i++{
        for j := i+1; j < n; j++{
            if min(height[i], height[j]) * (j-i) > maxNum {
                maxNum = min(height[i], height[j]) * (j-i)
            }
        }
    }
    return maxNum
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

//方法三：双指针，左右两边
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxNum := 0
    for left < right {
        area := 0
        if height[left] < height[right] {
            area = height[left]*(right-left)
            left++
        } else {
            area = height[right]*(right-left)
            right--
        }
        if area > maxNum {
            maxNum = area
        }
    }
    return maxNum
}