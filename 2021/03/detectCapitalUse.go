package main

//520. 检测大写字母

//模拟
func detectCapitalUse(word string) bool {
	cap, Ncap := 0, 0
	for _, ch := range word {
		if capJudge(ch) {
			cap++
		} else {
			Ncap++
		}
	}
	if cap == len(word) || Ncap == len(word) || cap == 1 && capJudge(rune(word[0])) {
		return true
	}
	return false

}

func capJudge(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}
