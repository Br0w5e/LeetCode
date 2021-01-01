package main

func isMagic(target []int) bool {

	nums := make([]int, len(target))
	for i := range nums {
		nums[i] = i+1
	}

	nums = helper(nums)
	// find k
	k := 0
	for k < len(nums) && nums[k] == target[k] {
		k++
	}
	if k == 0 {
		return false
	}

	for {
		if k >= len(nums) {
			return isSame(nums, target, len(nums))
		}
		if !isSame(nums, target, k) {
			return false
		}
		nums = nums[k:]
		target = target[k:]
		nums = helper(nums)
	}
	return false
}

func isSame(nums, target []int, end int) bool {
	for i := 0; i < end; i++ {
		if nums[i] != target[i] {
			return false
		}
	}
	return true
}

func helper(nums []int) []int {
	res := make([]int, len(nums))
	idx := 0
	for i := 1; i < len(nums); i+=2 {
		res[idx] = nums[i]
		idx++
	}
	for i := 0; i < len(nums); i+= 2 {
		res[idx] = nums[i]
		idx++
	}
	return res
}