package main

// 不是很懂，打卡题。

//685. 冗余连接 II
//并查集
func findRedundantDirectedConnection(edges [][]int) []int {
	numNodes := len(edges)
	uf := newUnionFind(numNodes+1)
	parent := make([]int, numNodes+1) //parent[i] 表示i的父节点
	for i, _ := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int

	//判断是否有一个节点存在两个父节点的情况
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) { // from 和 to 已经连接
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}
	//不存在节点两个父节点的情况，则附加边导致环
	if conflictEdge == nil {
		return cycleEdge
	}
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from int, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}

	return unionFind{ancestor}
}


//出入度，判断
var pre []int

func find(x int) int {
	if pre[x] != x {
		pre[x] = find(pre[x])
	}
	return pre[x]
}

var degree []int

func findRedundantDirectedConnection2(edges [][]int) []int {
	n := len(edges) + 1
	degree = make([]int, n)
	pre = make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = i
	}
	var e int // 记录入度为2的顶点
	for _, edge := range edges {
		v := edge[1]
		degree[v]++
		if degree[v] == 2 {
			e = v
			break
		}
	}

	tmp := []int{}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if v == e {
			// 不连接弧头为e的边，并将其保存下来
			tmp = append(tmp, u)
			continue
		}
		pu, pv := find(u), find(v)
		if pu == pv {
			// 不存在入度为2的顶点，
			// 存在某个指向根节点的顶点
			return []int{u, v}
		}
		pre[pu] = pv
	}

	// 如果某条边的弧尾与e位于同一棵子树下，则该边即为多余的边
	if find(tmp[0]) == find(e) {
		return []int{tmp[0], e}
	} else {
		return []int{tmp[1], e}
	}
}



