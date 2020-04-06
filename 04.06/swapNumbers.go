//不用中间变量，交换两个数字
//加法
func swapNumbers(numbers []int) []int {
    numbers[0] = numbers[0] + numbers[1]
    numbers[1] = numbers[0] - numbers[1]
    numbers[0] = numbers[0] - numbers[1]
    return numbers
}

//异或
func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[0] ^ numbers[1]
	numbers[0] = numbers[0] ^ numbers[1]

	return numbers
}