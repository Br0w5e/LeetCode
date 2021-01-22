package main

import (
	"math"
	"sort"
)
//并查集方法
type unionFind struct {
	parent []int
	rank []int
}

func minCostConnectPoints(points [][]int) int {
	res := 0
	n := len(points)
	type edge struct {
		v int
		w int
		dis int
	}
	edges := make([]edge, 0)
	for i, p := range points {
		for j := i+1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(p, points[j])})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})

	uf := newUnionFind(n)
	left := n-1
	for _, e := range edges {
		if uf.union(e.v, e.w) {
			res += e.dis
			left--
			if left == 0 {
				break
			}
		}
	}
	return res
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool{
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}

	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}

	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	return true
}

//Kruskal
func minCostConnectPoints(points [][]int) int {
	visited := make([]bool, len(points))
	minDist := make([]int, len(points))
	for i := 1; i < len(minDist); i++ {
		minDist[i] = math.MaxInt32
	}

	res := 0

	for _ = range visited {
		next := -1
		d := math.MaxInt32
		for j := range visited {
			if !visited[j] && minDist[j] < d {
				d = minDist[j]
				next = j
			}
		}
		res += d
		visited[next] = true

		for i := range visited {
			if !visited[i] {
				minDist[i] = min(minDist[i], getDistance(points[i], points[next]))
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getDistance(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}