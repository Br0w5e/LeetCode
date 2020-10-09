package main

import "fmt"

//834. 树中距离之和
//此题不会，打卡呢

func sumOfDistancesInTree(N int, edges [][]int) []int {
	res := make([]int, N)
	// edges form a tree, no cycle as compared to graph
	tree := make([][]int, N)

	// adjacent list to represent tree
	for i := range edges {
		parent, child := edges[i][0], edges[i][1]
		// undirected edge
		tree[parent] = append(tree[parent], child)
		tree[child] = append(tree[child], parent)
	}
	// count[i] means amount of nodes in the tree rooted at i-th node
	count := make([]int, N)
	var postOrder func(current, previous int)
	postOrder = func(current, previous int) {
		for _, child := range tree[current] {
			// in adjacent list, there could form a loop without recording visited
			if child == previous {
				continue
			}
			postOrder(child, current)
			// now count[i] is ready to use due to previous invocation
			count[current] += count[child]
			res[current] += res[child] + count[child]
		}
		count[current]++
	}
	var preOrder func(current, previous int)
	preOrder = func(current, previous int) {
		for _, child := range tree[current] {
			if child == previous {
				continue
			}
			// current is i's parent root
			res[child] = res[current] - count[child] + N - count[child]
			// log.Println(res)
			preOrder(child, current)
		}
	}
	// 1 <= N
	postOrder(0, -1)
	// now res[0] is determined
	preOrder(0, -1)
	// preOrder(0, -1, &res, count, tree, N)
	return res
}

func main()  {
	n := 6
	nums := [][]int{{0,1},{0,2},{2,3},{2,4},{2,5}}
	fmt.Println(sumOfDistancesInTree(n, nums))
}