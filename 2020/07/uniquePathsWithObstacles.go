//从原点达达目的地的路径
//第一行有1则置为0
//第一列有1则置为0
//其他dp[i][j] = dp[i-1][j] + dp[i][j-1]
package main
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    if n == 0 || m == 0 {
        return 0
    }
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++{
        dp[i] = make([]int, m)
        for j := 0; j < m; j++{
            if i == 0 && j == 0 {
                dp[i][j] = 1
            //第一行
            } else if i== 0 && j != 0{
                if obstacleGrid[i][j] == 0 {
                    dp[i][j] = dp[i][j-1]
                }
            //第一列
            } else if i != 0 && j == 0 {
                if obstacleGrid[i][j] == 0 {
                    dp[i][j]= dp[i-1][j]
                }
            } else {
                if obstacleGrid[i][j] == 0 {
                    dp[i][j] = dp[i-1][j]+dp[i][j-1]
                }
            }
        }
    }
    return dp[n-1][m-1]
}

//不添加额外空间
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    //开始位置就有障碍
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            //第一个元素
            if i == 0 && j == 0 {
                obstacleGrid[i][j] = 1
                //第一行
            } else if i == 0 && j != 0 {
                if obstacleGrid[i][j] == 1 {
                    obstacleGrid[i][j] = 0
                } else {
                    obstacleGrid[i][j] = obstacleGrid[i][j-1]
                }
                //第一列
            } else if i != 0 && j == 0 {
                if obstacleGrid[i][j] == 1{
                    obstacleGrid[i][j] = 0
                } else {
                    obstacleGrid[i][j] = obstacleGrid[i-1][j]
                }
                //其他
            } else {
                if obstacleGrid[i][j] == 0 {
                    obstacleGrid[i][j] = obstacleGrid[i-1][j]+obstacleGrid[i][j-1]
                } else {
                    obstacleGrid[i][j] = 0
                }
            }
        }
    }
    return obstacleGrid[n-1][m-1]
}