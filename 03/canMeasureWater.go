//判断两个给定的数字能否组成第三个数字，或者叫做倒水问题
//方法：裴蜀定理、扩展的欧几里得
func canMeasureWater(x int, y int, z int) bool {
	if x + y < z {
		return false
	}
	if x + y == z {
		return true
	}
	if z%gcd(x, y) != 0 {
		return false
	}
	return true
}

func gcd(a, b int) int {
	if a > b{
		a, b = b, a
	}
	for a != 0 {
		a, b = b%a, a
	}
	return b
}