package main

//395. 至少有K个重复字符的最长子串


//滑动窗口
//不是很懂
func longestSubstring(s string, k int) (int) {
	res := 0
	for t := 1; t <= 26; t++ {
		cnt := make([]int, 26)
		total := 0
		lessK := 0
		l := 0
		for r, ch := range s {
			ch -= 'a'
			if cnt[ch] == 0 {
				total++
				lessK++
			}
			cnt[ch]++
			if cnt[ch] == k {
				lessK--
			}

			for total > t {
				ch := s[l] - 'a'
				if cnt[ch] == k {
					lessK++
				}
				cnt[ch]--
				if cnt[ch] == 0 {
					total--
					lessK--
				}
				l++
			}
			if lessK == 0 {
				res = max(res, r-l+1)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
