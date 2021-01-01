//最后一块剩余的石头
//模拟
func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	size := len(stones)
	sort.Ints(stones)
	for stones[size-2] != 0 {
		stones[size-1] -= stones[size-2]
		stones[size-2] = 0
		sort.Ints(stones)
	}
	return stones[size-1]
}