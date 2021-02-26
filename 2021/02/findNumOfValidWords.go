package main

import "math/bits"

//暴力--->超时
func findNumOfValidWords(words []string, puzzles []string) []int {
	m := make(map[string][]string)
	for _, puzzle := range puzzles {
		puzzleMap := make(map[byte]bool, 0)
		for i := 0; i < len(puzzle); i++ {
			puzzleMap[puzzle[i]] = true
		}
		for _, word := range words {
			flag := false
			for i := 0; i < len(word); i++ {
				if word[i] == puzzle[0] {
					flag = true
					break
				}
			}
			if !flag {
				continue
			}
			j := 0
			for j = 0; j < len(word); j++ {
				if puzzleMap[word[j]] == false {
					j--
					break
				}
			}
			if j == len(word) {
				m[puzzle] = append(m[puzzle], puzzle)
			}
		}
	}
	res := make([]int, 0)
	for _, puzzle := range puzzles {
		res = append(res, len(m[puzzle]))
	}
	return res
}

//二进制状态压缩 pass
func findNumOfValidWords1(words []string, puzzles []string) []int {
	const puzzleLength = 7
	cnt := make(map[int]int, 0)
	for _, s := range words {
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(uint(mask)) <= puzzleLength {
			cnt[mask]++
		}
	}

	res := make([]int, len(puzzles))
	for i, s := range puzzles {
		first := 1 << (s[0] - 'a')
		// 枚举子集
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')
		}
		subset := mask
		for {
			res[i] += cnt[subset|first]
			subset = (subset - 1) & mask
			if subset == mask {
				break
			}
		}
	}
	return res
}
