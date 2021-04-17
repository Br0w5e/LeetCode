package main

//693. 交替位二进制数

//设置pre和cur
func hasAlternatingBits(n int) bool {
	pre := n&1
	for n != 0 {
		n /= 2
		cur := n&1
		//当前和上一个相同的时候返回false
		if cur == pre {
			return false
		}
		pre = cur
	}
	return true
}