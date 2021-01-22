package main

//989. 数组形式的整数加法

//数字的数数组形式相加
func addToArrayForm(A []int, K int) []int {
	numsK := make([]int, 0)
	for K != 0 {
		numsK = append(numsK, K%10)
		K /= 10
	}

	res := make([]int, max(len(A), len(numsK))+1)
	sum := 0
	numsK = reverse(numsK)

	for k, i, j := len(res)-1, len(A)-1, len(numsK)-1; i >= 0 || j >= 0; k, i, j = k-1, i-1, j-1 {
		if i < 0 {
			res[k] = (numsK[j] + sum) % 10
			sum = (numsK[j] + sum) / 10
		} else if j < 0 {
			res[k] = (A[i] + sum) % 10
			sum = (A[i] + sum) / 10
		} else {
			res[k] = (A[i] + numsK[j] + sum) % 10
			sum = (A[i] + numsK[j] + sum) / 10
		}
	}
	if sum != 0 {
		res[0] = sum
		return res
	}
	return res[1:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}
