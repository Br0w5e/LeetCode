func orangesRotting(grid [][]int) int {
    var dx = [4]int{0, 0, 1, -1}
    var dy = [4]int{-1, 1, 0, 0}
    m, n := len(grid), len(grid[0])
    
    fresh := 0
    rottens := make([][2]int, 0, m*n)
    //标记腐烂的坐标
    for i, row := range grid{
        for j, rank := range row{
            switch rank{
                case 1:
                    fresh++
                case 2:
                    rottens = append(rottens, [2]int{i, j})
            }
        }
    }

    //开始全部腐烂
    if fresh == 0{
        return 0
    }

    count := -1

    for len(rottens)>0 {
        count++
        size := len(rottens)
        //同时感染
        for r := 0; r < size; r++{
            x, y := rottens[r][0], rottens[r][1]
            //四个方向
            for k := 0; k < 4; k++{
                i, j := x+dx[k], y+dy[k]
                if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
                    continue
                }
                //在里边进行标记
                grid[i][j] = 2
                fresh--
                rottens = append(rottens, [2]int{i, j})
            }
        }
        rottens = rottens[size:]
    }
    if fresh > 0{
        return -1
    }
    return count
}