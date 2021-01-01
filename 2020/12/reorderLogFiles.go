package main


//937. 重新排列日志文件
import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		l1 := strings.Split(logs[i], " ")
		l2 := strings.Split(logs[j], " ")
		t1, t2 := 0, 0
		// 字母
		if l1[1][0] < '0' || l1[1][0] > '9'  {
			t1 = 1
		}
		// 字母
		if l2[1][0] < '0' || l2[1][0] > '9' {
			t2 = 1
		}
		// 字母日志 都排在 数字日志 之前。
		if t1 == 1 && t2 == 0 {
			return true
		}
		if t1 == 0 && t2 == 1 {
			return false
		}
		// 数字日志 应该按原来的顺序排列。
		if t1 == 0 && t2 == 0 {
			return i < j
		}
		for i := 1; i < len(l1) || i < len(l2); i++ {
			if i < len(l1) && i < len(l2) {
				if l1[i] != l2[i] {
					return l1[i] < l2[i]
				}
				continue
			}
			if i < len(l1) {
				return true
			}
			if i < len(l2) {
				return false
			}
		}
		// 字母日志 在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序；
		return l1[0] < l2[0]
	})
	return logs
}
