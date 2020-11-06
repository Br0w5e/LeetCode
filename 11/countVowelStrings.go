package main

//5555. 统计字典序元音字符串的数目
func countVowelStrings(n int) int {
	nums := []int{1, 1, 1, 1, 1}
	sum := 5
	for i := 2; i <= n; i++ {
		sum = 1
		for j := 1; j < 5; j++ {
			nums[j] += nums[j-1]
			sum += nums[j]
		}
	}
	return sum
}