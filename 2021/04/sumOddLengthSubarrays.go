package main

import "fmt"

//1588. 所有奇数长度子数组的和
func sumOddLengthSubarrays(arr []int) int {
	res := 0
	n := len(arr)
	//每次递增2
	for g := 1; g < n+1; g = g + 2 {
		//从头开始
		for i := 0; i <= n-g; i++ {
			//求里边的和
			for j := i; j < i+g; j++ {
				res += arr[j]
			}
		}
	}
	return res
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))
}
