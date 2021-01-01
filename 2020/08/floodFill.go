//图像填充
//DFS
//参考岛屿面积和坏的的橙子
package main
var dx = [4]int{0, 0, 1, -1}
var dy = [4]int{1, -1, 0, 0}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    color := image[sr][sc]
    if newColor != color {
        dfs(image, sr, sc, newColor, color)
    }
    return image
}

func dfs(image [][]int, sr int, sc int, newColor int, color int) {
    if sr < 0 || sr == len(image) || sc < 0 || sc == len(image[0]) || image[sr][sc] != color {
		return
	}
    image[sr][sc] = newColor
    for k := 0; k < 4; k++{
		dfs(image, sr+dx[k], sc+dy[k], newColor, color)
	}
}