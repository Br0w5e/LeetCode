package main

//1053. 交换一次的先前排列

//倒序遍历
//第一步：从当前序列的后往前找，找到第一个降序的位置（A[i]>A[i+1]），则必存在能构造比当前小的序列。
//
//第二步：把A[i]后面的数字中，小于A[i]且最接近A[i]的值的数字找出来，和A[i]交换。
//
//第一步不再往前找，因为往前找更换，会让小的值出现在高位，导致不是最大字典序。

func prevPermOpt1(A []int) []int {
	for i := len(A)-2; i >= 0; i-- {
		if A[i] > A[i+1] { //此处存在逆序，需要移动A[i]
			flag := i+1   //标记
			max := -1
			for j := i+2; j < len(A); j++ { //寻找替换位置
				if A[j] < A[i] && A[j] > max { // 必须满足 A[j] < A[i] 并且最接近A[i]，否则不能满足交换后的字典序小于原始字典序
					max = A[j]
					flag = j
				}
			}
			A[i], A[flag] = A[flag], A[i]
			return A
		}
	}
	return A
}
