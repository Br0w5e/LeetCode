package main

//796. 旋转字符串

func rotateString(A string, B string) bool {
	if  len(A) == 0 && len(B) == 0 {
		return true
	}
	//旋转n次
	for i := 0; i < len(A); i++ {
		A = A[1:] + A[:1]
		if A == B {
			return true
		}
	}
	return false
}