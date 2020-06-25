//三色排序
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	p := 0
	for p <= right {
		if nums[p] == 0{
			nums[left], nums[p] = nums[p], nums[left]
			left++
			p++
		} else if nums[p] == 1 {
			p++
		} else if nums[p] == 2 {
			nums[right], nums[p] = nums[p], nums[right]
			//还需要进行下一轮判断，所以没有p--
			right--
		}
	}
	return 
}