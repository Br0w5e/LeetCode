package main

//5620. 连接连续二进制数字
import (
	"fmt"
	"math/bits"
)

func concatenatedBinary(n int) int {
	mod := int(1e9)+7
	res := 0
	x := 1
	for i := n; i >= 1; i-- {
		for j := i; j > 0; j, x = j/2, x*2%mod {
			if j %2 == 1 {
				res = ((res%mod) + (x%mod))%mod
			}
		}
	}
	return res
}
func concatenatedBinary2(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res = (res<<bits.Len(uint(i)) | i) % (1e9 + 7)
	}
	return res
}

func main()  {
	n := 13
	fmt.Println(concatenatedBinary(n))
}