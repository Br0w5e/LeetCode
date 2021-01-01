package main

import (
	"fmt"
	"sort"
)

//重新安排行程
//欧拉图 半欧拉图

func findItinerary(tickets [][]string) []string {
	var (
		m = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		//源点为key 生成m
		m[src] = append(m[src], dst)
	}
	for key := range m {
		//排序
		sort.Strings(m[key])
	}
	//递归
	var dfs func(curr string)
	dfs = func(curr string) {
		for len(m[curr]) > 0 {
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}
	//更好立理解的写法
	//dfs = func(curr string) {
	//	for {
	//		//k不存在，或者k存在但没有下一个元素（终点）
	//		if v, ok := m[curr]; !ok || len(v) == 0 {
	//			break
	//		}
	//		tmp := m[curr][0]
	//		m[curr] = m[curr][1:]
	//		dfs(tmp)
	//	}
	//	res = append(res, curr)
	//}
	dfs("JFK")
	//反向输出
	for left, right := 0, len(res)-1 ; left < right; left, right = left+1, right-1{
		res[left], res[right] = res[right], res[left]
	}
	return res

}

func main()  {
	tickets := [][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}
	fmt.Println(findItinerary(tickets))
}