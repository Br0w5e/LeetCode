package main
//1018. 可被 5 整除的二进制前缀

//模拟，注意越界的情况
func prefixesDivBy5(A []int) []bool {
	num := 0
	res := make([]bool, len(A))
	for i := 0; i < len(A); i++ {
		num = num*2 + A[i]
		if num % 5 == 0 {
			res[i] = true
		}
		num %= 5
	}
	return res
}
