package main

//997. 找到小镇的法官

//抽象为图
func findJudge(N int, trust [][]int) int {
	//入度表示被信任
	in := make([]int, N+1)
	//出度表示信任
	out := make([]int, N+1)

	for _, t := range trust {
		in[t[1]]++
		out[t[0]]++
	}
	//寻找入度为N-1，出度为0的人
	for i := 1; i <= N; i++ {
		if in[i] == N-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
