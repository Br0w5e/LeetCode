//两人平分糖果，判断一人能够分到的最大种类数
//map
func distributeCandies(candies []int) int {
    m := make(map[int]int)
    for _, candy := range candies {
        m[candy]++
    }
    return min(len(candies)/2, len(m))
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}