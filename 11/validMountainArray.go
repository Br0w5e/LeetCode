package main

//941. 有效的山脉数组

func validMountainArray(A []int) bool {
	index := getMax(A)
	if index == 0 || index == len(A)-1 {
		return false
	}
	for i := 0; i < index; i++ {
		if A[i] >= A[i+1] {
			return false
		}
	}
	for i := len(A)-1; i > index; i-- {
		if A[i-1] <= A[i] {
			return false
		}
	}
	return true
}

func getMax(nums []int) int {
	left, right := 0, 0
	for left, right = 0, len(nums)-1; left < right; {
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return left
}

func validMountainArray(a []int) bool {
	i, n := 0, len(a)

	// 递增扫描
	for ; i+1 < n && a[i] < a[i+1]; i++ {
	}

	// 最高点不能是数组的第一个位置或最后一个位置
	if i == 0 || i == n-1 {
		return false
	}

	// 递减扫描
	for ; i+1 < n && a[i] > a[i+1]; i++ {
	}

	return i == n-1
}
