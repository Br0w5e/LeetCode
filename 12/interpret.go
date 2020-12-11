package main
//5617. 设计 Goal 解析器
//模拟
func interpret(command string) string {
	res := ""
	for i := 0; i < len(command); {
		if command[i] == 'G' {
			res += "G"
			i++
		} else if i < len(command)-1 && command[i:i+2] == "()" {
			res += "o"
			i += 2
		}  else if i < len(command)-3 && command[i:i+4] == "(al)" {
			res += "al"
			i+=4
		}
	}
	return res
}
