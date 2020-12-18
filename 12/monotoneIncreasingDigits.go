package main

//738. 单调递增的数字
func monotoneIncreasingDigits(n int) int {
	res, tmp, base := n%10, 9, 10
	for n/10 != 0 {
		if n/10%10 > n %10 {
			res = (n/10%10-1)*base + tmp
			n = n/10-1
		} else {
			res = n/10%10*base + res
			n /= 10
		}
		tmp = tmp*10+9
		base *= 10
	}
	return res
}

func main()  {
	monotoneIncreasingDigits(371)
}