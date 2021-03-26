package main

import (
	"fmt"
	"strings"
)

//求1法，在当前位置为1的时候加1
func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if (1<<i & num) > 0 {
			res++
		}
	}
	return res
}



//化0法，每次将最低位1去掉
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res++
		//将最低位转化为1
		num &= (num-1)
	}
	return res
}

//一行超人
func hammingWeight1(num uint32) int {
	return strings.Count(fmt.Sprintf("%b", num), "1")
}

