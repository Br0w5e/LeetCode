package main

import "sort"

//621. 任务调度器
// A->X->X->A->X->X->A，这里的X表示除了A以外其他字母，或者是待命，不用关心具体是什么，反正用来填充两个A的间隔的。上面执行顺序的规律是： 有count - 1个A，其中每个A需要搭配n个X，再加上最后一个A，所以总时间为 (count - 1) * (n + 1) + 1
// 要注意可能会出现多个频率相同且都是最高的任务，比如 ["A","A","A","B","B","B","C","C"]，所以最后会剩下一个A和一个B，因此最后要加上频率最高的不同任务的个数 maxCount
// 公式算出的值可能会比数组的长度小，如["A","A","B","B"]，n = 0，此时要取数组的长度
//map建表
func leastInterval(tasks []byte, n int) int {
	//建表
	m := make(map[byte]int)
	for _, task := range tasks {
		m[task]++
	}
	//寻找出现次数最大的
	maxCount := 0
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}
	//次数最大的出现次数
	flag := 0
	for _, v := range m {
		if v == maxCount {
			flag++
		}
	}
	//最小值是len(tasks)
	return max((maxCount-1)*(n+1)+flag, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//数组统计字频
func leastInterval2(tasks []byte, n int) int {
	//统计字频
	count := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		count[tasks[i]-'A']++
	}
	//排序，最后一个肯定是最大的
	sort.Ints(count)
	//计算有多少个最大的
	maxCount := 0
	for i := 25; i >= 0; i-- {
		if count[i] != count[25] {
			break
		}
		maxCount++
	}
	return max((count[25] - 1) * (n + 1) + maxCount, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}