package main

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	index := findIndex(mountainArr, 0, length-1)
	if mountainArr.get(index) == target {
		return index
	}
	index = findLeftIndex(mountainArr, 0, index-1, target)
	if index != -1 {
		return index
	}
	return findRightIndex(mountainArr, index+1, length-1, target)
}
//找山顶
func findIndex(mountainArray *MountainArray, left int, right int)  int {
	for left < right {
		mid := left + (right-left)/2
		if mountainArray.get(mid) < mountainArray.get(mid+1) {
			left = mid+1
		} else {
			right = mid
		}
	}
	return left
}
//升序
func findLeftIndex(mountainArray *MountainArray, left int, right int, target int) int {
	for left < right {
		mid := left + (right-left)/2
		if mountainArray.get(mid) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if mountainArray.get(left) == target {
		return left
	}
	return -1
}
//降序
func findRightIndex(mountainArray *MountainArray, left int, right int, target int) int {
	for left < right {
		mid := left + (right-left+1)/2
		if mountainArray.get(mid) < target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if mountainArray.get(left) == target {
		return left
	}
	return -1
}
