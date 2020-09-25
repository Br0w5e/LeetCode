//寻找最小的，一般方法
package main

//153 寻找旋转数组中的最小值
func findMin(numbers []int) int {
	i := 0
	if len(numbers) == 1 {
		return numbers[0]
	}
	if len(numbers) == 2 {
		if numbers[0] > numbers[1] {
			return numbers[1]
		} else {
			return numbers[0]
		}
	} else {
		for i < len(numbers)-2 {
			if numbers[i] <= numbers[i+1] {
				i++
			} else {
				return numbers[i+1]
			}
		}
		if numbers[len(numbers)-2] > numbers[len(numbers)-1] {
			return numbers[len(numbers)-1]
		} else {
			return numbers[0]
		}

	}
}

func findMin2(nums []int) int {
	low := 0
	high := len(nums) - 1
	for nums[low] >= nums[high] && low < high {
		mid := (low + high) / 2
		if nums[mid] < nums[low] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			low++
		}
	}
	return nums[low]
}
