//计算给定数字的各位积和的差
//最直观的思路就是将每个数字提取出来放在数组中，最后求和求积求差
func subtractProductAndSum(n int) int {
    res := make([]int, 0)
    for n != 0 {
        res = append(res, n%10)
        n = n/10
    }
    sum, mul := 0, 1
    for _, num := range res{
        sum += num
        mul *= num
    }
    return mul-sum
}
//改进一下
func subtractProductAndSum(n int) int{
	sum, mul := 0, 1
	for n != 0{
		sum += (n%10)
		mul *= (n%10)
		n /= 10
	}
	return mul-sum
}