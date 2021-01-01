package main

//806. 写字符串需要的行数
//模拟
func numberOfLines(widths []int, S string) []int {
	row, col := 1, 0
	for i := 0; i < len(S); i++ {
		if col + widths[int(S[i]-'a')] <= 100 {
			col += widths[int(S[i]-'a')]
		} else {
			row++
			col = widths[int(S[i]-'a')]
		}
	}
	return []int{row, col}
}
