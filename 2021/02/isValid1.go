package main
//591. 标签验证器

// 标签验证器
// 思路：就是单纯的栈 中间会设计到很多的判断 0ms/2.1mb
func isValid(code string) bool {
	n := len(code)
	if n < 2 || code[0] != '<' || code[1] == '!'{
		return false
	}

	tagStack := make([]string,0)
	s := ""
	for key,v := range code {

		// 遇到特殊处理将其内容置空
		if len(s) >= 11 && s[:9] == "<![CDATA[" && s[len(s)-2:] == "]]" && v == '>'{
			s = ""
			continue
		}

		if v == '<' && s == ""{
			s = "<"
			continue
		}

		if v == '>' {
			if s == ""{
				continue
			}
			s += ">"
			if checkTag(s) {
				if s[0:2] == "</" {
					if len(tagStack) > 0 && tagStack[len(tagStack)-1] == s[2:len(s)-1]{
						tagStack = tagStack[:len(tagStack)-1]
						s = ""
						// 开头 只能被 最后 干掉
						if key != n-1 && len(tagStack) == 0{
							return false
						}
						continue
					} else {
						return false
					}
				} else if s[0:2] != "<!"{
					tagStack = append(tagStack,s[1:len(s)-1])
					s = ""
					continue
				}
			}else {
				return false
			}
		}

		if len(s) > 0{
			s += string(v)
		}
	}

	if s == "" && len(tagStack) == 0{
		return true
	}
	return false
}

func checkTag(s string) bool {
	if len(s) >= 9 && s[:9] == "<![CDATA["{
		return true
	}
	if len(s) <= 2 {
		return false
	}

	count := 0
	for key,val := range s[1:len(s)-1] {
		if val == '/' || key == 0{
			continue
		}
		if val < 65 || val >90 {
			return false
		}
		count++
	}
	if count > 9{
		return false
	}

	return true
}
