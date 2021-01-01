//返回a, b中的最大值 不用判断语句
func maximum(a, b int) int{
	flag := uint64(a-b) >> 63
	// a > b, flag = 0, return a
	// a < b, flag = 1; return 1
	return int(flag)*b + int(1^flag)*a
}