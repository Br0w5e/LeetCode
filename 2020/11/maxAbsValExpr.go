package main

//1131. 绝对值表达式的最大值
//暴力破解，超时
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	max := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1); j++ {
			if abs(arr1[i], arr1[j]) + abs(arr2[i], arr2[j]) + abs(i, j) > max{
				max = abs(arr1[i], arr1[j]) + abs(arr2[i], arr2[j]) + abs(i, j)
			}
		}
	}
	return max
}

func abs(a, b int) int {
	if a > b {
		return a-b
	}
	return b-a
}

//化解为数学问题求解
func maxAbsValExpr2(arr1 []int, arr2 []int) int {
	res := 0
	n := make([]int, 4)
	n[0] = arr1[0]+arr2[0]
	n[1] = -arr1[0]-arr2[0]
	n[2] = arr1[0]-arr2[0]
	n[3] = -arr1[0]+arr2[0]
	for i := 1; i < len(arr1); i++ {
		res = max(res, n[0]-arr1[i]-arr2[i]+i)
		n[0] = max(n[0],arr1[i]+arr2[i]-i)
		res = max(res,n[1]+arr1[i]+arr2[i]+i)
		n[1] = max(n[1],-arr1[i]-arr2[i]-i)
		res = max(res,n[2]-arr1[i]+arr2[i]+i)
		n[2] = max(n[2],arr1[i]-arr2[i]-i)
		res = max(res,n[3]+arr1[i]-arr2[i]+i)
		n[3] = max(n[3],-arr1[i]+arr2[i]-i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}