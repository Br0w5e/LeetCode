package main

import "sort"

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefix := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	l, r, ans := 0, arr[n-1], -1
	for l <= r {
		mid := (l + r) / 2
		index := sort.SearchInts(arr, mid)
		if index < 0 {
			index = -1 * index - 1
		}
		cur := prefix[index] + (n - index) * mid
		if cur <= target {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	chooseSmall := check(arr, ans)
	chooseBig := check(arr, ans + 1)
	if abs(chooseSmall - target) > abs(chooseBig - target) {
		ans++
	}
	return ans
}

func check(arr []int, x int) int {
	ret := 0
	for _, num := range arr {
		ret += min(num, x)
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefix := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	r := arr[n-1]
	ans, diff := 0, target
	for i := 1; i <= r; i++ {
		index := sort.SearchInts(arr, i)
		if index < 0 {
			index = -index - 1
		}
		cur := prefix[index] + (n - index) * i
		if abs(cur - target) < diff {
			ans = i
			diff = abs(cur - target)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}