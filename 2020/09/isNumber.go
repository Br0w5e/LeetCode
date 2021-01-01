package main

import (
	"regexp"
	"strings"
)

//正则
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	reg := "^[+|-]?(([0-9]*\\.?[0-9]+)|([0-9]+\\.?[0-9]*)|[0-9]+)((e|E)[+|-]?[0-9]+)?$"
	res, _ := regexp.MatchString(reg, s)
	return res
}

//纯粹
func isNumber2(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	res := isInt(&s)
	if len(s) > 0 && s[0] == '.' {
		s = s[1:]
		res = isUInt(&s) || res
	}

	if len(s) > 0 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		res = res && isInt(&s)
	}
	return len(s) == 0 && res
}

func isUInt(s *string) bool {
	i := 0
	for ; i < len(*s) && (*s)[i] >= '0' && (*s)[i] <= '9'; i++ {
	}
	*s = (*s)[i:]
	return i > 0
}

func isInt(s *string) bool {
	if len(*s) > 0 && ((*s)[0] == '-' || (*s)[0] == '+') {
		*s = (*s)[1:]
	}
	return isUInt(s)
}
