package main

//406. 根据身高重建队列
import "sort"

func reconstructQueue(people [][]int) [][]int {
	//身高升序、排位降序
	sort.SliceStable(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})
	//结果
	res := make([][]int, len(people))
	for _, person := range people {
		//前边的个数
		cnt := person[1]
		//插入就行了
		for i, _ := range res {
			if res[i] == nil {
				if cnt == 0 {
					res[i] = person
					break
				}
				cnt--
			}
		}
	}
	return res
}