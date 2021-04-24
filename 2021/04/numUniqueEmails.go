package main

//929. 独特的电子邮件地址
import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, email := range emails {
		emailArr := strings.Split(email, "@")
		//域名不变
		local, domain := emailArr[0], emailArr[1]
		//处理本地
		for i := 0; i < len(local); i++ {
			if local[i] == '.' {
				local = local[:i] + local[i+1:]
			} else if local[i] == '+' {
				local= local[:i]
				break
			}
		}
		m[local+"@"+domain]++
	}
	return len(m)
}

func main()  {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
