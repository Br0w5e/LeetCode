package main

//781. 森林中的兔子

// https://leetcode-cn.com/problems/rabbits-in-forest/solution/sen-lin-zhong-de-tu-zi-by-leetcode-solut-kvla/
//贪心，看官方题解
func numRabbits(answers []int) int {
	m := make(map[int]int)
	res := 0
	for _, answer := range answers {
		m[answer]++
	}
	for k, v := range m {
		res += (k + v) / (k + 1) * (k + 1)
	}
	return res
}
