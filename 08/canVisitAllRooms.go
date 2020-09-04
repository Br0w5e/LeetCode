package main

//dfs
var (
	num int
	visit []bool
)
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	visit = make([]bool, n)
	//从第一个房子开始
	dfs(rooms, 0)
	return num == n
}

func dfs(rooms [][]int, x int) {
	//标记为访问过的
	visit[x] = true
	num++
	//x中含有的钥匙
	for _, key := range rooms[x] {
		//该房子没被访问
		if !visit[key] {
			dfs(rooms, key)
		}
	}
}

// bfs
func canVisitAllRooms2(rooms [][]int) bool {
	n := len(rooms)
	num := 0
	visit := make([]bool, n)
	//开始时候队列里只有0号房间
	queue := make([]int, 1)
	visit[0] = true

	for i := 0; i < len(queue); i++ {
		x := queue[i]
		num++
		//判断队列中的房子中的钥匙，以及相关的房子是否被打开
		for _, key := range rooms[x] {
			//含有钥匙未打开
			if !visit[key] {
				//可访问
				visit[key] = true
				//用来操作后边节点
				queue = append(queue, key)
			}
		}
	}
	return num == n
}