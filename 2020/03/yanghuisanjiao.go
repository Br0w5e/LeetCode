//生成杨辉三角
func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++{
		single_res := make([]int, i+1)
		res = append(res, single_res)
		//开头一个
		res[i][0] = 1
		for j := 1; j < i; j++{
			res [i][j] = res[i-1][j-1] + res[i-1][j]
		}
		//最后一个
		res[i][i] = 1
	}
	return res
	
}

//返回杨辉三角的第N行
//修改上面算法
func getRow(rowIndex int) []int {
    res := make([][]int, 0)
    for i := 0; i < rowIndex+1; i++{
        single_res := make([]int, i+1)
        res = append(res, single_res)
        res[i][0] = 1
        for j := 1; j < i; j++ {
            res[i][j] = res[i-1][j-1] + res[i-1][j]
        }
        res[i][i] = 1
    }
    return res[rowIndex]
}


//
func getRow(rowIndex int) []int {
	rowIndex++
	res := make([]int, rowIndex)
	if len(res) < 1 {
		return res
	}

	res[0] = 1
	if len(res) == 1 {
		return res
	}
	
	res[1] = 1
	if len(res) == 2 {
		return res
	}

	for i := 2; i < rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}
