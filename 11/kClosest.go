package main

import "sort"

//973. 最接近原点的 K 个点
//暴力
func kClosest(points [][]int, K int) [][]int {
	m := make([]int, 0)
	t := make([]int, 0)
	res := make([][]int, 0)
	for _, point := range points {
		m = append(m, distance(point))
		t = append(t, distance(point))
	}
	sort.Ints(t)
	for i := 0; i < K; i++ {
		for j := 0; j < len(m); j++ {
			if m[j] == t[i] {
				res = append(res, points[j])
			}
		}
		if len(res) == K {
			return res
		}
	}
	return res
}

func distance(nums []int) int {
	return nums[0]*nums[0]+nums[1]*nums[1]
}

//
func kClosest(points [][]int, K int) [][]int {
	//稳定的
	sort.SliceStable(points, func(i, j int) bool {
		p, q := points[i], points[j]
		//目标排序方式
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:K]
}