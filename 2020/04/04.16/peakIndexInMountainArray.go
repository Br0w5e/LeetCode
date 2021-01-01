//寻找最高峰的位置
//二分法，注意结束位置
func peakIndexInMountainArray(A []int) int {
	left, right := 0, len(A)-1
	for left <= right {
		mid := (left+right)/2
		if A[mid] > A[mid+1] {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return left
}