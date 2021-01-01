//通法
func addDigits(num int) int {
	for num >= 10 {
		next := 0
		for num != 0 {
			next += num%10
			num /= 10
		}
		num = next
	}
	return num
}

//骚操作
func addDigits(num int) int {
    return (num-1)%9 + 1
}