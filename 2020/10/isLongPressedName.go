package main

//925. 长按键入

//双指针，进行操作，联系子串匹配
func isLongPressedName(name string, typed string) bool {
	index1, index2 := 0, 0
	for index2 < len(typed) {
		//防止index1 越界
		if index1 < len(name) && name[index1] == typed[index2] {
			index1++
			index2++
			//第一个字符判断
		} else if index2 > 0 && typed[index2] == typed[index2-1] {
			index2++
		} else {
			return false
		}
	}
	return index1 == len(name)
}

