//计算平方根
//方法1：二分查找法
func mySqrt(x int) int {
	left := 1
	right := x
	for left <= right {
		mid := (left + right) >> 1
		if mid * mid > x{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int(right)
}

//方法二： 牛顿迭代法
func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) >> 1
	}
	return r
}