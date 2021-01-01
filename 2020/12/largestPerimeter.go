package main

//976. 三角形的最大周长
import (
	"fmt"
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A)-1; i >= 2; i-- {
		if A[i-1] + A[i-2] > A[i] {
			return A[i]+A[i-1]+A[i-2]
		}
	}
	return 0
}

func main()  {
	A := []int{2,1,2}
	fmt.Println(largestPerimeter(A))
}
