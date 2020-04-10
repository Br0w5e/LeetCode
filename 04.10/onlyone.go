//只出现一次的数字
//总思路：异或运算

//类型一：整个数组中只有一个数字出现了一次，其余均出现两次
func singleNum(nums []int) int {
	single := 0
	for _, num := range nums{
		single ^= num
	}
	return single
}

//类型二：除了一个数字出现一次，其他都出现了三次
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}  
//map,计数法别忘了
func singleNumber(nums []int) int {
	count := make(map[int]int)
	for _,v := range nums{
			count[v]++
	}
	for k,v := range count{
		if v == 1{
			return k
		}
	}
	return 0
}

//类型三：数组中有两个数字仅出现一次，其余均出现两次
//思路：利用异或结果分组，再将分组结果进行找出
func singleNumbers(nums []int) []int{
	or := 0
	for _, num := range nums {
		or ^= num
	}
	//寻找最低位是一的数
	right := 1
    for or & right == 0 {
        right <<= 1
    }
    onlyOne := 0
    //两组中分别处理
	for _, num := range nums{
		if num & right != 0 {
			onlyOne ^= num
		}
	}
	return []int{onlyOne, or^onlyOne}
}