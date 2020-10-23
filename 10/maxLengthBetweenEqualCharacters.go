package main

//5543. 两个相同字符之间的最长子字符串
//注意和219 containsNearbyDuplicate 的区别， 当相同字母出现的时候并不需要改变字典里边记录的值
func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune]int)
	max := -1
	for i, ch := range s {
		if pos, ok := m[ch]; ok {
			//判断距离是不是大于最大距离
			if i-pos-1 > max {
				max = i - pos - 1
			}
		} else {
			//如果字符没出现过，在字典里边进行记录
			m[ch] = i
		}
	}
	return max
}