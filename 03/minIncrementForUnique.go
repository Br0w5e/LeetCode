//求最小增量。使得数组中每个元素出现的次数唯一。
//阶梯法，砌台阶
func minIncrementForUnique(A []int) int {
	if len(A) == 0{
		return 0
	}
	sort.Ints(A)
	res, value := 0, A[0]
	for i := 1; i < len(A); i++ {
		if A[i] <= value {
			value++
			res += value - A[i]
		} else {
			value = A[i]
		}
	}
	return res
}