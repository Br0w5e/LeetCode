package main

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	rows, columns := len(matrix), len(matrix[0])
	sorted := make([]int, rows * columns)
	index := 0
	for _, row := range matrix {
		for _, num := range row {
			sorted[index] = num
			index++
		}
	}
	sort.Ints(sorted)
	return sorted[k-1]
}
