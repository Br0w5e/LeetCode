package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for k, _ := range numSet {
		if !numSet[k-1] {
			currentNum := k
			currentSreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentSreak++
			}
			if longestStreak < currentSreak {
				longestStreak = currentSreak
			}
		}
	}
	return longestStreak
}

func main()  {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("%d", longestConsecutive(nums))
}