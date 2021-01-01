package main

import "sort"

//1356. 根据数字二进制下 1 的数目排序
//朴素的算法
func sortByBits(arr []int) []int {
	sort.Ints(arr)
	max := 0
	m := make(map[int][]int)
	//以1的个数为键建立哈希表， 并记录最大1的个数
	for _, num := range arr {
		cnt := count(num)
		if cnt > max {
			max = cnt
		}
		m[cnt] = append(m[cnt], num)
	}
	//遍历哈希表
	res := make([]int, 0)
	for i := 0; i <= max; i++ {
		for k := 0; k < len(m[i]); k++ {
			res = append(res, m[i][k])
		}
	}
	return res
}

func count(n int) int {
	res :=0
	for n != 0 {
		res += n%2
		n /= 2
	}
	return res
}

//调库
func sortByBits2(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool{
		cntx, cnty := count(arr[i]), count(arr[j])
		return  cntx < cnty || (cntx == cnty && arr[i] < arr[j])
	})
	return arr
}