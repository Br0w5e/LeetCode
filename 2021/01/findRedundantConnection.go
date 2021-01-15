package main

//684. 冗余连接

//并查集
type UnionFindSet struct {
	parent []int
	count int
}

func (u *UnionFindSet) init(l int)  {
	u.parent = make([]int, l)
	for i := 0; i < l; i++ {
		u.parent[i] = i
	}
	u.count = len(u.parent)
}

func (u *UnionFindSet) union(i, j int) bool {
	a, b := u.find(i), u.find(j)
	if a == b {
		return true
	}
	u.parent[a] = b
	return false
}

func (u *UnionFindSet) find(i int) int {
	for i != u.parent[i] {
		i = u.parent[i]
	}
	return i

}
func findRedundantConnection(edges [][]int) []int {
	res := [2]int{}
	u := &UnionFindSet{}
	u.init(len(edges)+1)
	for _, v := range edges {
		if u.union(v[0], v[1]) {
			res[0] = v[0]
			res[1] = v[1]
		}
	}
	return res

}
