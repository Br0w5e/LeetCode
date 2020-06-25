//使得數組中數字相等的最小操作，每次操作n-1個

//方法：反向思維，每次操作一次， 寻找最小数字

func minMoves(nums []int) int {
	min := getMin(nums)
	res := 0
	for _, num := range nums{
		res += (num-min)
	}
	return res
}

func getMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] > nums[right]{
			left++
		} else {
			right--
		}
	}
	return nums[left]
}


func minMoves(nums []int) int {
	l, res := 0, 0
	min := math.MaxInt32
	for _, num := range nums{
		res += num
		l++
		if num < min {
			min = num
		}
	}
	return res - min*l
}



//使得数组中的数字相等的最小操作，每次只能操作一个，加一或者减一
//方法 寻找中位数
func minMoves2(nums []int) int {
    sort.Ints(nums)
    left, right := 0, len(nums)-1
    res := 0 
    for left < right {
        res += (nums[right]-nums[left])
        right--
        left++
    }
    return res
}