package main

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := sLen+1
	ansL, ansR := -1, -1
	check := func() bool{
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r-l+1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func minWindow2(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen == 0 || tLen == 0 || sLen < tLen {
		return ""
	}
	var i, j, start, found int
	minLen, tCount, sCount := sLen+1, [256]int{}, [256]int{}
	for _, c := range t {
		tCount[c]++
	}
	for j < sLen {
		if found < tLen {
			prev := int(s[j])
			if tCount[prev] > 0 {
				sCount[prev]++
				if sCount[prev] <= tCount[prev] {
					found++
				}
			}
			j++
		}
		for found == tLen {
			if j-i < minLen {
				start, minLen = i, j-i
			}
			next := int(s[i])
			if tCount[next] > 0 {
				sCount[next]--
				if sCount[next] < tCount[next] {
					found--
				}
			}
			i++
		}
	}
	if minLen == sLen+1 {
		return ""
	}
	return s[start : start+minLen]
}