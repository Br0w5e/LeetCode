package main

//1017. 负二进制转换
func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	base := -2
	res := make([]byte, 0)
	for N != 0 {
		r := N % base
		N /= base
		if r < 0 {
			r += (-base)
			N++
		}
		res = append(res, byte(r+int('0')))
	}
	res = reverse(res)
	return string(res)
}

func reverse(nums []byte) []byte {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
