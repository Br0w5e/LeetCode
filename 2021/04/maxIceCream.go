package main

//5735. 雪糕的最大数量

//排序，先买价值小的
import "sort"

func maxIceCream(costs []int, coins int) int {
	res := 0
	sort.Ints(costs)
	for _, cost := range costs {
		if cost > coins {
			break
		}
		res++
		coins -= cost
	}
	return res
}
