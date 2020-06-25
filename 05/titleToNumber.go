package main

import "math"

func titleToNumber(s string) int {
	res := 0
	for _, ch := range s {
		num := int(byte(ch)-'A')+1
		res = res * 26 + num
	}
	return  res
}

func titleToNumber2(s string) int {
	res, n := 0, len(s)
	for i := 0; i < n; i++{
		res += int(math.Pow(26, float64(n-i-1))) * (int(s[i]-'A')+1)
	}
	return res
}