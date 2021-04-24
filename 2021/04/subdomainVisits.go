package main

//811. 子域名访问计数

//map 记录
import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, domain := range cpdomains {
		domainArr := strings.Split(domain, " ")
		cnt, _ := strconv.Atoi(domainArr[0])
		m[domainArr[1]] += cnt
		for i := 0; i < len(domainArr[1])-1; i++ {
			if domainArr[1][i] == '.' {
				m[domainArr[1][i+1:]] += cnt
			}
		}
	}
	res := make([]string, 0)
	for k, v := range m {
		res = append(res, strconv.Itoa(v) + " " + k)
	}
	return res
}
