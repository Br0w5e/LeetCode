package main

import (
	"math"
	"sort"
)

type unionFind struct {
	parent, size []int
	setCount     int //当前连通分量
}
//并查集，Kruskal算法
func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
	return true
}

//计算最小生成树权值
func calcMST(uf *unionFind, ignoreID int, edges [][]int) int {
	mstValue := 0
	for i, e := range edges {
		if i != ignoreID && uf.union(e[0], e[1]) {
			mstValue += e[2]
		}
	}
	if uf.setCount > 1 {
		return math.MaxInt64
	}
	return mstValue
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	//边标号
	for i, e := range edges {
		edges[i] = append(e, i)
	}
	//边安权值排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	//计算最小生成树价值
	mstValue := calcMST(newUnionFind(n), -1, edges)

	var keyEdges, pseudokeyEdges []int
	//枚举所有边
	for i, e := range edges {
		// 是否为关建边，忽略该边，得到的最小生成树权值大于不忽略该边时
		if calcMST(newUnionFind(n), i, edges) > mstValue {
			keyEdges = append(keyEdges, e[3])
			continue
		}
		// 是否为伪关键边，忽略该边，得到的最小生成树加上忽略边权值恰好等于最小生成树权值
		uf := newUnionFind(n)
		uf.union(e[0], e[1])
		if e[2]+calcMST(uf, i, edges) == mstValue {
			pseudokeyEdges = append(pseudokeyEdges, e[3])
		}
	}
	return [][]int{keyEdges, pseudokeyEdges}
}
