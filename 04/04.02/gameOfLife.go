func gameOfLife(board [][]int){
	var dx = [8]int{0, 0, 1, 1, 1, -1, -1, -1}
	var dy = [8]int{1, -1, 0, 1, -1, 0, 1, -1}

	n := len(board)
	m := len(board[0])
	if n == 0 {
		return
	}

	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
            //计算周围活着的数量
			cnt := 0
			for k := 0; k < 8; k++{
				x := i + dx[k]
				y := j + dy[k]
				if x >= 0 && x < n && y >= 0 && y < m {
					cnt += board[x][y] & 1
				}  
			}
            //准备进行操作
			if board[i][j] > 0 {
				//活的
				if cnt >= 2 && cnt <= 3 {
                    //仍然活着
					board[i][j] = 0b11
				}
                //其他都死了，因为是 0b01所以不用操作
			} else if cnt == 3 {
                //死了复活
				board[i][j] = 0b10
			}
		}
	}
    //去掉最后一位更新状态
	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			board[i][j] >>= 1
		}
	}
	return
}