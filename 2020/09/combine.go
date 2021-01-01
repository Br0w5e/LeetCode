//返回n以内，k个数的所有组合
package main

import "fmt"

// 77 组合
var result [][]int

func combine(n int, k int) [][]int {
	result = [][]int{}
	if n == 0 || k == 0 {
		return result
	}

	nums := []int{}
	process(1, n, k, nums)
	return result
}

func process(start, n, k int, nums []int) {
	if len(nums) == k {
		tmp := make([]int, k)
		copy(tmp, nums)
		result = append(result, tmp)
		return
	}

	for i := start; i <= n-k+len(nums)+1; i++ {
		nums = append(nums, i)
		process(i+1, n, k, nums)
		nums = nums[0 : len(nums)-1]
	}

}

func combine2(n int, k int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	for i := 0; i < k; i++ {
		tmp = append(tmp, i+1)
	}
	tmp = append(tmp, n+1)

	for i := 0; i < k; {
		comb := make([]int, k)
		copy(comb, tmp[:k])
		res = append(res, comb)
		for i = 0; i < k && tmp[i]+1 == tmp[i+1]; i++ {
			tmp[i] = i + 1
		}
		tmp[i]++
	}
	return res
}

//暴力组合
func combine3(n int, k int) [][]int {
	res := make([][]int, 1)
	//获取每个元素
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	for _, num := range nums {
		for _, ele := range res {
			tmp := make([]int, len(ele), len(ele)+1)
			copy(tmp, ele)
			tmp = append(tmp, num)
			res = append(res, tmp)
		}
	}
	ret := make([][]int, 0)
	for _, v := range res {
		if len(v) == k {
			ret = append(ret, v)
		}
	}
	return ret
}

func main() {
	var str string
	var a int
	fmt.Scan(&str, &a)
	fmt.Println(str, a)
}
