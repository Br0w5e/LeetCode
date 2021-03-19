package main

// 54. 螺旋矩阵

//順时针打印二维数组
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }

    step := 0
    size := len(matrix)*len(matrix[0])
    nums := make([]int, size)
    //端点
    top, bottom, right, left := 0, len(matrix)-1, len(matrix[0])-1, 0
    for step < size {
        //左到右
        for i := left; i <= right && step < size; i++{
            nums[step] = matrix[top][i]
            step++
        }
        top++
        //上到下
        for i := top; i <= bottom && step < size; i++{
            nums[step] = matrix[i][right]
            step++
        }
        right--

        //右到左
        for i := right; i >= left && step < size; i-- {
            nums[step] = matrix[bottom][i]
            step++
        }
        bottom--
        //下到上
        for i := bottom; i >= top && step < size; i-- {
            nums[step] = matrix[i][left]
            step++
        }
        left++

    }

    return nums
}

//帶有方向控制的
func spiralOrder(matrix [][]int) []int {
    var dx = [4]int{0, 1, 0, -1}
    var dy = [4]int{1, 0, -1, 0}
    m := len(matrix)
    if m == 0 {
        return []int{}
    }
    n := len(matrix[0])
    res := make([]int, 0)
    x, y, dw := 0, 0, 0
    //标记矩阵
    grid := make([][]int, m)
    for i, _ := range grid {
        grid[i] = make([]int, n)
    }
    for i := 0; i < m*n; i++{
        grid[x][y] = 1
        res = append(res, matrix[x][y])
        if  x+dx[dw] < 0 || x+dx[dw] >= m || y+dy[dw] < 0 || y+dy[dw] >= n || grid[x+dx[dw]][y+dy[dw]] != 0 {
            dw = (dw+1)%4
        }
        x, y = x+dx[dw], y+dy[dw]
    }
    return res
}
