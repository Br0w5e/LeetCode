//判断图像重叠面积
//只有是1的时候才算重叠
//方法：暴力破解，都试一遍。
func largestOverlap(A [][]int, B [][]int) int {
    n := len(A)
    maxoverlap := 0
    overlap := 0
    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            //B不动
            overlap = check(A, B, i, j, n)
            if overlap > maxoverlap {
                maxoverlap = overlap
            }
            //A不动
            overlap = check(B, A, i, j, n)
            if overlap > maxoverlap {
                maxoverlap = overlap
            }
        }
    }
    return maxoverlap
}

func check(A [][]int, B [][]int, x int, y int, size int) int {
    ret := 0
    lenX := size - x
    lenY := size - y
    for i := 0; i <  lenX; i++{
        for j := 0; j < lenY; j++{
            if B[i][j] == 1 && A[x+i][y+j] == 1{
                ret++
            }
        }
    }
    return ret
}