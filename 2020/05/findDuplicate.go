package main

func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] == 1{
			return num
		} else {
			m[num]++
		}
	}
	return -1
}

func findDuplicate1(nums []int) int {
	slow := nums[0]
	fast := nums[slow]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = nums[slow]
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func findDuplicate2(nums []int) int {
	n := len(nums)
	ans := 0
	bit_max := 31
	for ((n-1) >> bit_max) == 0 {
		bit_max--
	}
	for bit := 0; bit <= bit_max; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if (nums[i] & (1 << bit)) > 0 {
				x++
			}
			if i >= 1 && (i & (1 << bit)) > 0 {
				y++
			}
		}
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}

func findDuplicate3(nums []int) int {
	n := len(nums)
	l, r := 1, n - 1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}