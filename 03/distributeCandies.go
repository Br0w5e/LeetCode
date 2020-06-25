// 对用户进行分糖果
func distributeCandies(candies int, num_people int) []int {
	res := make([]int , num_people)
	for i := 0; candies > 0; i++{
		res[i % num_people] += min(i+1, candies)
		candies -= (i+1)
	}
	return res
}

func min(a, b int) int {
	if a > b{
		return b
	}
	return a
}