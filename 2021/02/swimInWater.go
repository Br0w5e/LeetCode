package main

import "sort"

//778. 水位上升的泳池中游泳
//参考 1631. 最小体力消耗路径

//二分法
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) (ans int) {
	n := len(grid)
	return sort.Search(n*n-1, func(threshold int) bool {
		if threshold < grid[0][0] {
			return false
		}
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < n && !vis[x][y] && grid[x][y] <= threshold {
					vis[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return vis[n-1][n-1]
	})
}



//并查集
type unionFind struct {
	parent, size []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) (ans int) {
	n := len(grid)
	pos := make([]pair, n*n)
	for i, row := range grid {
		for j, h := range row {
			pos[h] = pair{i, j} // 存储每个平台高度对应的位置
		}
	}

	uf := newUnionFind(n * n)
	for threshold := 0; ; threshold++ {
		p := pos[threshold]
		for _, d := range dirs {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < n && grid[x][y] <= threshold {
				uf.union(x*n+y, p.x*n+p.y)
			}
		}
		if uf.inSameSet(0, n*n-1) {
			return threshold
		}
	}
}