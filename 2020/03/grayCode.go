//格雷码生成
//总结规律

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	size := int(math.Exp2(float64(n)))
	data := make([]int, size)
	data[0] = 0
	base := 1
	for digit := 0; digit < n; digit++{
		for i := 0; i < base; i++ {
			data[base+i] = data[base-i-1] + base
		}
		base *=2
	}
	return data
}
//移位操作
func grayCode(n int) []int {
	ret := make([]int, 1<<n)
	for i := 0; i < n; i++ {
		for j, k := 0, (1<<(i+1))-1; j < k; j, k = j+1, k-1 {
			ret[k] = ret[j] + (1 << i)
		}
	}
	return ret
}