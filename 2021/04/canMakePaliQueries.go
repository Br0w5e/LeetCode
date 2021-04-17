package main


//1177. 构建回文串检测
func canMakePaliQueries(s string, queries [][]int) []bool {
	cnt := []int{0}
	for i, v := range s {
		cnt = append(cnt, cnt[i]^1<<(int(v)-97))
	}
	ans := []bool{}
	for _, v := range queries {
		i, j, k := v[0], v[1], v[2]
		ans = append(ans, func(num int) (res int) {
			for num > 0 {
				res += num & 1
				num >>= 1
			}
			return
		}(cnt[i]^cnt[j+1])/2 <= k)
	}
	return ans
}
