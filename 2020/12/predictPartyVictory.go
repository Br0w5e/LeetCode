package main

//649. Dota2 参议院

//循环队列
func predictPartyVictory(senate string) string {
	radiant := make([]int, 0)
	dire := make([]int, 0)
	for i, s := range senate {
		//将元素位置加入队列
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	//循环至某一队列为空
	for len(radiant) > 0 && len(dire) > 0 {
		//小的元素可以禁用大的元素，并可以参加下一轮投票
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
