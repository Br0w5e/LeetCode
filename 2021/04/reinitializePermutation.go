package main

//1806. 还原排列的最少操作步数

//模拟
func reinitializePermutation(n int) int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	arr := make([]int, n)
	res := 0
	//完成一次操作
	for !judge(arr) {
		res++
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		copy(perm, arr)
	}

	return res
}

func judge(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return false
		}
	}
	return true
}
