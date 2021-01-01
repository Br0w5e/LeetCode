package main

import "fmt"

//并查集
func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}
	//扫描等式
	for _, str := range equations {
		if str[1] == '=' {
			index1 := int(str[0]-'a')
			index2 := int(str[3]-'a')
			union(parent, index1, index2)
		}
	}
	//扫描不等式
	for _, str := range equations {
		if str[1] == '!' {
			index1 := int(str[0]-'a')
			index2 := int(str[3]-'a')
			if find(parent, index1) == find(parent, index2) {
				return false
			}
		}
	}
	return true
}
//并
func union(parent []int, index1 int, index2 int) {
	parent[find(parent, index1)] = find(parent, index2)
}
//查
func find(parent []int, index int) int{
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}

func main()  {
	equations := []string{"c==c","b==d","x!=z"}
	fmt.Println(equationsPossible(equations))
}
