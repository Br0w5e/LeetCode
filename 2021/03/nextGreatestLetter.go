package main


//744. 寻找比目标字母大的最小字母

//循环出现很重要
func nextGreatestLetter(letters []byte, target byte) byte {
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}
