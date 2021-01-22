package main
//721. 账户合并

//哈希表+并查集
import "sort"

func accountsMerge(accounts [][]string) [][]string {
	emailIndex := make(map[string]int)
	emailName := make(map[string]string)
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailIndex[email]; !has {
				emailIndex[email] = len(emailIndex)
				emailName[email] = name
			}
		}
	}

	parent := make([]int, len(emailIndex))
	for i := range parent {
		parent[i] = i
	}

	var find func( int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, account := range accounts {
		firstIdx := emailIndex[account[1]]
		for _, email := range account[2:] {
			union(emailIndex[email], firstIdx)
		}
	}
	res := make([][]string, 0)
	indexEmail := make(map[int][]string)
	for email, index := range emailIndex {
		index = find(index)
		indexEmail[index] = append(indexEmail[index], email)
	}

	for _, emails := range indexEmail {
		sort.Strings(emails)
		account := append([]string{emailName[emails[0]]}, emails...)
		res = append(res, account)
	}
	return res
}