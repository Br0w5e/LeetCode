package main

//386. 字典序排数
//迭代
func lexicalOrder(n int) []int {
	res := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		res[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			if cur >= n {
				cur /= 10
			}
			//加数
			cur++
			//退位
			for cur%10 == 0 {
				cur /= 10
			}
		}
	}
	return res
}

//dfs
func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	for i := 1;i <= 9;i++ {
		dfs(i, n, &res)
	}
	return res

}

func dfs(dg, n int, res *[]int) {
	if dg > n {
		return
	}
	*res = append(*res, dg)
	for i := 0;i <=9;i++ {
		dfs(10 * dg + i, n, res)
	}
}