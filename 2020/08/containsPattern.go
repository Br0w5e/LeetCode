package main

import "fmt"

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i < len(arr)-m+1; i++ {
		flag := 0
		pattern := arr[i:i+m]
		for j := i; j < len(arr)-m+1; j += m {
			if judge(arr[j:j+m], pattern) {
				flag++
			} else {
				break
			}
			if flag == k {
				return true
			}
		}
	}
	return false
}

func judge(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main()  {
	arr := []int{2, 2}
	m := 1
	k := 2
	fmt.Println(containsPattern(arr, m, k))
}