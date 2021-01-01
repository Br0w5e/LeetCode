//三色排序
package main
//75. 颜色分类

//left指向0， right 指向2， p指向1
func sortColors(nums []int)  {
	left, right := 0, len(nums)-1
	p := 0
	for p <= right {
		switch nums[p] {
		case 0:
			nums[left], nums[p] = num[p], num[left]
			p++
			left++
		case 1:
			p++
			//跟0交换后，继续需要跟1交换
		case 2:
			nums[right], nums[p] = nums[p], nums[right]
			right--
		}
	}
	return
}

//两轮
//第一轮：将0都放到最左边
//第二轮：将2都放到最右边
func sortColors2(nums []int) {
	for left, right := 0, len(nums)-1; left < right; {
		if nums[right] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		} else {
			right--
		}
	}
	for left, right := 0, len(nums)-1; left < right; {
		if nums[left] == 2 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return
}