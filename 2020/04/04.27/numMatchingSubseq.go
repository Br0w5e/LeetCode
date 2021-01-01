package main
func numMatchingSubseq(S string, words []string) int {
	res := 0
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	for k, v := range m{
		if judge(S, k) {
			res += v
		}
	}
	return res
}

func judge(text string, sub string) bool {
	j, n := 0, len(sub)
	for i := 0; i < len(text) && j < n; i++{
		if text[i] == sub[j] {
			j++
		}
		if j == n {
			return true
		}
	}
	return false
}