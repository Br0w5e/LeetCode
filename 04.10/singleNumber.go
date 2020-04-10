//寻找出现一次的数字
//方法1：map
func singleNumber(nums []int) int {
    res := make(map[int]int)
    for _, num := range nums {
        res[num]++
    }
    for k, v := range res {
        if v == 1 {
            return k
        }
    }
    return 0
}

//方法二：参考onlyone.go  也就是位运算
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}