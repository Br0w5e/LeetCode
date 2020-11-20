package main

//907. 子数组的最小值之和
import "math"
//暴力
func sumSubarrayMins(arr []int) int {
	MOD := 1e9 + 7
	res := 0
	for i := 1; i <= len(arr); i++ {
		for j := 0; j < len(arr)+1-i; j++ {
			tmp := make([]int, i)
			copy(tmp, arr[j:j+i])
			res += findMin(tmp)
		}
	}
	return res % int(MOD)
}

func findMin(nums []int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

//单调栈
func sumSubarrayMins2(A []int) int {
	result := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	dp := make([]int, len(A)+1)
	mod := int(math.Pow(10, 9)) + 7
	for index := 0; index < len(A); index++ {
		for stack[len(stack)-1] != -1 && A[index] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[index+1] = dp[stack[len(stack)-1]+1] + (index-stack[len(stack)-1])*A[index]
		result = (result + dp[index+1]) % mod
		stack = append(stack, index)
	}
	return result
}
