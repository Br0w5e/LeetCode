package main

//1128. 等价多米诺骨牌对的数量
//暴力-->超时
func numEquivDominoPairs(dominoes [][]int) int {
	res := 0
	for i := 0; i < len(dominoes); i++ {
		for j := i+1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || (dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
				res++
			}
		}
	}
	return res
}

//map 标记
//超时一般选择map标记
func numEquivDominoPairs2(dominoes [][]int) int {
	res := 0
	m := make(map[int]int)
	for _, dominoe := range dominoes {
		//每个有序
		if dominoe[0] > dominoe[1] {
			dominoe[0], dominoe[1] = dominoe[1], dominoe[0]
		}
		//1 <= dominoes.length <= 40000
		//1 <= dominoes[i][j] <= 9
		//等价牌有相同坐标
		mark := dominoe[0]*10+dominoe[1]
		//与之前的每一个都可以组成对
		res += m[mark]
		//标记加一
		m[mark]++
	}
	return res
}