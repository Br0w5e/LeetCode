package main

//795. 区间子数组个数
import "fmt"

//区间大于等于L小于等于R的子数组个数 = 区间小于等于R的子数组个数 - 区间小于L的子数组个数
// 也 = 区间小于等于R的子数组个数 - 区间小于等于L-1的子数组个数；
func numSubarrayBoundedMax(A []int, L int, R int) int {
	return help(A, R) - help(A, L-1)
}

func help(A []int, max int) int {
	res := 0
	numSubarray := 0
	for _, num := range A {
		if num <= max {
			//每次在上一个基础上加一
			numSubarray++
			//结果得在上一个基础上面加，想想排列组合
			res += numSubarray
		} else {
			numSubarray = 0
		}
	}
	return res
}


func main()  {
	A := []int{2,1,4,3}
	L := 2
	R := 3

	fmt.Println(numSubarrayBoundedMax(A, L, R))
}