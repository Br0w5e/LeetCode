package main

//1268. 搜索推荐系统
import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := make([][]string, 0)
	for i := 1; i <= len(searchWord); i++ {
		prex := searchWord[:i]
		res = append(res, make([]string, 0))
		for _, product := range products {
			//字符串本身比product小
			if len(product) < i {
				continue
			}
			//相同加入
			if  prex == product[:i] {
				res[i-1] = append(res[i-1], product)
			}
			//三个足够
			if len(res[i-1]) == 3 {
				break
			}
		}
	}
	return res
}