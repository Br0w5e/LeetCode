//判断有效括号
func isValid(s string) bool {
	var c []byte
	symbol := map[byte]byte{
		')' : '(',
		']' : '[',
		'}' : '{',
	}
	for _, value := range s {
		n := len(c)
		if n > 0 {
            if c[n-1] == symbol[byte(value)] {
                c = c[:n-1]
                continue
			}
		}
		c = append(c, byte(value))
	}
	return len(c) == 0
}