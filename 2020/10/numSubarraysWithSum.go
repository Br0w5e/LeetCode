package main

//前缀和
//P[i] = A[0] + A[1] + ... + A[i - 1]
//P[j + 1] - P[i] = A[i] + A[i + 1] + ... + A[j]，  快速计算A[i……j]的和
func numSubarraysWithSum(A []int, S int) int {
	prefix := make(map[int]int)
	//res 存储最终结果
	//sum记录当前前缀和
	res, sum := 0, 0
	for _, a := range A {
		// 累加，用于计算前缀和
		sum += a
		if sum == S {
			res++
		}
		// A[0:i] - A[0:j] == S  ;
		//即 A[i:j]和等于S prefix存储了符合条件的数量
		res += prefix[sum-S]
		// 存储前缀和
		prefix[sum]++
	}
	return res
}

//改进版本
func numSubarraysWithSum1(A []int, S int) int {
	prefix := make(map[int]int)
	//和为0不选就可以
	prefix[0] = 1
	res, sum := 0, 0
	for _, a := range A {
		sum += a
		res += prefix[sum-S]
		prefix[sum]++
	}
	return res
}
