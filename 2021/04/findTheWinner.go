package main

import "fmt"

//5727. 找出游戏的获胜者

//模拟
func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	flag := 0
	i := 0
	for {
		tmp := 0
		for ; tmp < k; i++ {
			if nums[i%n] != 0 {
				tmp++
			}
		}
		flag++
		nums[(i-1)%n] = 0
		if flag == n-1 {
			break
		}
	}
	for _, num := range nums {
		if num != 0 {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println(findTheWinner(5, 2))
}
