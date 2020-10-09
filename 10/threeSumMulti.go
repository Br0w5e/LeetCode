package main
//923. 三数之和的多种可能
//三数之和的多种可能

func threeSumMulti(A []int, target int) int {
	count := 0
	mod := 1000000007
	nums := [101]int{}
	for _, a := range A {
		nums[a]++
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				continue
			}
			k := target - i - j
			if k < j {
				break
			}
			if k > 100 || nums[k] == 0 {
				continue
			}
			count = (count + cal(i, j, k, nums)) % mod
		}
	}
	return count
}

func cal(i, j, k int, nums [101]int) int {
	if i == j && j == k {
		num := nums[i]
		return num * (num-1) * (num-2) / (1 * 2 * 3)
	}
	if i == j && j != k {
		return nums[i] * (nums[i] - 1) * nums[k] / 2
	}
	if i != j && j == k {
		return nums[i] * nums[j] * (nums[j] - 1) / 2
	}
	return nums[i] * nums[j] * nums[k]
}