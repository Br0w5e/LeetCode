package main

//5734. 判断句子是否为全字母句

//map即可
func checkIfPangram(sentence string) bool {
	m := make([]int, 26)
	for _, w := range sentence {
		m[int(w-'a')] = 1
	}
	for _, v := range m {
		if v != 1 {
			return false
		}
	}
	return true
}
