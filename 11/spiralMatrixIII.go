package main

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	var next = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := make([][]int, 0)
	total := R * C
	k := 0
	x := r0
	y := c0
	for total > 0 {
		for i := 0; i < 4; i++ {
			if i == 0 || i == 2 {
				k++
			}
			for j := 0; j < k; j++ {
				if x < R && y < C && x >= 0 && y >= 0 {
					ans = append(ans, []int{x, y})
					total--
				}
				x = x + next[i][0]
				y = y + next[i][1]
			}
		}
	}
	return ans
}

func spiralMatrixIII2(R int, C int, r0 int, c0 int) [][]int {
	//1东  2南  3西  4北
	direct := 1
	length := 0
	result := [][]int{}
	last1, last2, last3, last4 := 0, 0, 0, 0
	length1, length2, length3, length4 := 1, 1, 2, 2
	if r0 < R && r0 >= 0 && c0 < C && c0 >= 0 {
		result = append(result, []int{r0, c0})
	}
	for len(result) < R*C {
		if direct == 1 {
			length1 += last1
			length = length1
			last1 = 2
		} else if direct == 2 {
			length2 += last2
			length = length2
			last2 = 2
		} else if direct == 3 {
			length3 += last3
			length = length3
			last3 = 2
		} else {
			length4 += last4
			length = length4
			last4 = 2
		}
		for temp := length; temp > 0; temp-- {
			if direct == 1 {
				c0++
			} else if direct == 2 {
				r0++
			} else if direct == 3 {
				c0--
			} else {
				r0--
			}
			if r0 < R && r0 >= 0 && c0 < C && c0 >= 0 {
				result = append(result, []int{r0, c0})
			}
		}
		direct++
		if direct == 5 {
			direct = 1
		}
	}
	return result
}

