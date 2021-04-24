package main

//面试题 05.01. 插入
import "fmt"

func insertBits(N int, M int, i int, j int) int {
	N_j := N >> (j + 1) << (j + 1)
	N_i := (N >> i << i) ^ N
	N_i_j := N_i | N_j
	return N_i_j | (M << i)
}

func main()  {
	N := 1024
	M := 19
	i := 2
	j := 6
	fmt.Println(insertBits(N, M, i, j))
}
