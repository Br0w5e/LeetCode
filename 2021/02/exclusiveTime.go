package main

//636. 函数的独占时间
import (
	"strconv"
	"strings"
)

//栈，符合后进先出，栈里存储函数编号
func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	pre, _ := strconv.Atoi(strings.Split(logs[0], ":")[2])
	stack := make([]int, 0)
	for _, log := range logs {
		split := strings.Split(log, ":")
		id, _ := strconv.Atoi(split[0])
		t, _ := strconv.Atoi(split[2])
		if split[1] == "start" {
			if len(stack) > 0 {
				res[stack[len(stack)-1]] += t - pre
			}
			stack = append(stack, id)
			pre = t
		} else {
			res[stack[len(stack)-1]] += t - pre + 1
			stack = stack[:len(stack)-1]
			pre = t + 1
		}
	}
	return res
}
