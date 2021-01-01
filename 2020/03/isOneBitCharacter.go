//判断最后一个字符是不是1比特字符。 10\11均为两比特、0为一比特
//方法一：从前往后遍历：遇见1跳两个，遇见0跳一个直到最后。 判断最后一个是不是独立出来的。
func isOneBitCharacter(bits []int) bool {
    n := len(bits) - 1
    i := 0 
    for i < n {
        i += bits[i] + 1
    }
    return i == n
}

//方法二：从后往前遍历，判断最后一位前边连1的个数
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	index := n - 2//倒数第二位
	cnt := 0
	for index >= 0 && bits[index] != 0 {
		cnt++
		index--
	}
	if cnt % 2 == 0{
		return true
	}
	return false
}