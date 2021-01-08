package main

//830. 较大分组的位置
import "fmt"

func largeGroupPositions(s string) [][]int {
	if len(s) < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	start, end := 0, 0
	for end < len(s){
		for end < len(s) && s[start] == s[end]{
			end++
		}
		if end-start > 2 {
			res = append(res, []int{start, end-1})
		}
		start = end
	}
	return res
}

func main()  {
	s := "abbxxxxzzy"
	fmt.Println(largeGroupPositions(s))
}
