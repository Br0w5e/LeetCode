package main

//969. 煎饼排序

import "fmt"


//例如:[3,2,4,1]---->[?,?,?,4]
//先找到数字4的位置,将数字4前进行翻转变成[4,2,3,1],接下来我们在整体翻转[1,3,2,4],这样把数字4移动列表底.
//然后,我们[1,3,2,4]--->[?,?,3,4],还是用刚才方法,首先找到数字3,翻转数字3前面的,再翻转已经排好数字(这里指数字4)前就可以了.


func pancakeSort(arr []int) []int {
	res := make([]int, 0)
	n := len(arr)
	for n > 0 {
		idx := getIndex(arr, n)
		res = append(res, idx+1)
		arr = pancake(arr, idx)
		res = append(res, n)
		arr = pancake(arr, n-1)
		n--
	}
	return res
}

func getIndex(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func pancake(nums []int, right int) []int {
	left := 0
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

func main()  {
	arr := []int{3,2,4,1}
	fmt.Println(pancakeSort(arr))
}