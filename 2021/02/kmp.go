package main

import "fmt"

//普通字符串匹配
func Sindex(S,T string) int {
	i := 0
	j := 0
	//同时满足才可以  找除字符串出现的第一个位置
	for ;i <= len(S) -1  && j <= len(T) -1;{
		if S[i] == T[j]{
			//当字符匹配时 i j 都加1
			i++
			j++
		}else{
			//不匹配 j 回溯到第一个字符 为 0   i 回溯 从下一个字符开始匹配
			i = i -j +1
			j = 0
		}
	}
	//如果 j 大于 或者 等于 T串的长度 说明匹配成功
	if j >= len(T) -1 {
		return i - len(T) +1
	}

	return 0
}

//next数组 获取
func get_next(T string) []int {
	i, j := 0, -1
	next := make([]int, 20)
	next[0] = -1
	for i<len(T)-1{
		if j == -1 || T[i] == T[j]{
			j++
			i++
			if T[i] == T[j] {
				next[i] = next[j]
			}else{
				next[i] = j

			}
		}else{
			j = next[j]
		}

	}
	return next
}

//func get_next(T string) [20]int {
//	j := -1
//	var next [20]int
//	next[0] = j;
//	for i := 1; i < len(T); i++{ // 注意i从1开始
//		for j >= 0 && T[i] != T[j + 1] { // 前后缀不相同了
//			j = next[j] // 向前回溯
//		}
//		if T[i] == T[j + 1] { // 找到相同的前后缀
//			j++
//		}
//		next[i] = j // 将j（前缀的长度）赋给next[i]
//	}
//	return next
//}

//KMP字符串匹配
func SindexKMP(S,T string) int {
	next := get_next(T)
	fmt.Println(next)
	i := 0
	j := 0
	//同时满足才可以  找除字符串出现的第一个位置
	for ;i <= len(S) -1  && j <= len(T) -1;{

		if j == -1|| S[i] == T[j]{
			//当字符匹配时 i j 都加1
			i++
			j++
		}else{
			//子串的 偏移量 从next数组中取  i 不变
			j = next[j]
		}
	}
	//如果 j 大于 或者 等于 T串的长度 说明匹配成功
	if j >= len(T) -1 {
		return i - len(T) +1
	}

	return 0
}

func main()  {
	S := "BBC ABCDAB ABCDABCDABDE"
	T := "ABCDABD"
	fmt.Println(SindexKMP(S, T))
}


