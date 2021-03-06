package main

//5689. 统计匹配检索规则的物品数量

//switch语句
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	for _, item := range items {
		switch ruleKey{
		case "type":
			if item[0] == ruleValue {
				res++
			}
		case "color":
			if item[1] == ruleValue {
				res++
			}
		case "name":
			if item[2] == ruleValue {
				res++
			}
		}
	}
	return res
}
