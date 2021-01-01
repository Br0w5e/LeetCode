package main

//旋转数组，最后一个是小于等于第一个的
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left+(right-left)/2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid+1
		} else {
			right--
		}
	}
	return numbers[left]
}


func minArray1(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}