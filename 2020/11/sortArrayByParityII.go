package main

//922. 按奇偶排序数组 II

//分别记录，然后重新赋值
func sortArrayByParityII(A []int) []int {
	m1 := make([]int, 0)
	m2 := make([]int, 0)
	for _, num := range A {
		if num%2 == 0 {
			m1 = append(m1, num)
		} else {
			m2 = append(m2, num)
		}
	}
	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			A[i] = m1[i/2]
		} else {
			A[i] = m2[i/2]
		}
	}
	return A
}

//双指针
func sortArrayByParityII2(A []int) []int {
	for i := 0; i < len(A); i++ {
		if (i+A[i])%2 != 0 {
			for j := i+1; j < len(A); j++ {
				if (i+A[j]) % 2 == 0 {
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		}
	}
	return A
}
