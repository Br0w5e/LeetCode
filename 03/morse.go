//验证摩尔斯电码
//思路：字典
func uniqueMorseRepresentations(words []string) int {
    morseStr := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	morseMap := map[string]int{}
	for _, word := range words {
		str := ""
		for _, char := range word {
			str += morseStr[char - 'a']
		}
		morseMap[str]++
	}
	return len(morseMap)
}