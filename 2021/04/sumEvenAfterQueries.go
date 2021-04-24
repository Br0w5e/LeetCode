package main

//985. 查询后的偶数和

//暴力
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, 0)
	for _, querie := range queries {
		sum := 0
		A[querie[1]] += querie[0]
		for _, a := range A {
			if a % 2  == 0  {
				sum += a
			}
		}
		res = append(res, sum)
	}
	return res
}
