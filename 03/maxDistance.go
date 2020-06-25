//计算最大距离
// 广度优先遍历
type point struct {
	x, y int
}
func maxDistance(grid [][]int) int {
	var queue []point
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				queue = append(queue, point{i, j})
			}
		}
	}

	rows := len(grid)
	cols := len(grid[0])
	if len(queue) == 0 || len(queue) == rows * cols {
		return -1
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dist := 0
	for len(queue) > 0 {
		var tempQueue []point
		for _, p := range queue {
			for _, dir := range dirs {
				newX := p.x + dir[0]
				newY := p.y + dir[1]
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && grid[newX][newY] == 0 {
					tempQueue = append(tempQueue, point{newX, newY})
					grid[newX][newY] = 1
				}
			}
		}
		dist++
		queue = tempQueue
	}
	return dist - 1
}

//深度优先遍历
func maxDistance(grid [][]int) int {
    max := -1
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            //陆地
            if grid[i][j] == 1 {
                continue
            }
            res := dfs(grid, i, j)
            if res > max {
                max = res
            }
        }
    }
    return max
}

func dfs(grid [][]int, x int, y int) int {
    length := 2*(len(grid)-1)
    for i := 1; i <= length; i++{
        for k := -1*i; k <= i; k++{
            r := x + k
            if r >= 0 && r < len(grid) {
                c1 := i - abs(k)
                c2 := abs(k) - i
                if c1 + y >= 0 && c1 + y < len(grid){
                    if grid[r][c1+y] == 1 {
                        return i
                    }
                }
                if c2 + y >= 0 && c2 + y < len(grid) {
                    if grid[r][c2+y] == 1 {
                        return i
                    }
                }
            }
        }
    }
    return -1
}

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}