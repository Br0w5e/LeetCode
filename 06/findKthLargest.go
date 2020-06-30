package main

func findKthLargest(nums []int, k int) int {
	for i := k / 2; i >= 0; i-- {
		minHeapify(nums[:k], i)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			minHeapify(nums[:k],0)
		}
	}
	return nums[0]
}

func minHeapify(nums []int, root int) {
	left, right := 2*root+1, 2*root+2
	min := root
	if left < len(nums) && nums[left] < nums[min] {
		min = left
	}
	if right < len(nums) && nums[right] < nums[min] {
		min = right
	}
	if min != root {
		nums[root], nums[min] = nums[min], nums[root]
		minHeapify(nums, min)
	}
}