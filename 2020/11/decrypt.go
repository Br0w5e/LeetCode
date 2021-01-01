package main

import "fmt"

//5550. 拆炸弹
func decrypt(code []int, k int) []int {
	tmp := make([]int, 2*len(code))
	for i := 0; i < len(tmp); i++ {
		if i < len(code) {
			tmp[i] = code[i]
		} else {
			tmp[i] = code[i-len(code)]
		}
	}
	res := make([]int, len(code))
	if k == 0 {
		return res
	} else if k > 0 {
		for i := 0; i < len(res); i++ {
			sum := 0
			for j := i+1; j < i+1+k; j++ {
				sum += tmp[j]
			}
			res[i] = sum
		}
	} else if k < 0{
		for i := 0; i < len(res); i++ {
			sum := 0
			for j := len(code)+i-1; j > len(code)+i-1+k; j-- {
				sum += tmp[j]
			}
			res[i] = sum
		}
	}
	return res
}

func main()  {
	code := []int{2,4,9,3}
	k := -2
	fmt.Println(decrypt(code, k))
}
