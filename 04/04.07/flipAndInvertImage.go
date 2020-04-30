//翻转图像
//先对折，再翻转：0->1,1->0
func flipAndInvertImage(A [][]int) [][]int {
    n := len(A)
    for i := 0; i < n; i++{
        for j := 0; j < (n+1)/2; j++{
            A[i][j], A[i][n-1-j] = A[i][n-1-j]^1, A[i][j]^1

        }
    }
    return A
}