package main

//1405. 最长快乐字符串
import (
	"bytes"
	"sort"
)

type Pair struct {
	c   rune
	cnt int
}

func longestDiverseString(a int, b int, c int) string {
	arr := []*Pair{
		{c: 'a', cnt: a},
		{c: 'b', cnt: b},
		{c: 'c', cnt: c},
	}
	var buf bytes.Buffer
	var l1, l2 rune
	for {
		// 根据数量升序
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].cnt < arr[j].cnt
		})
		var p *Pair
		// 贪心：最多连续放置2个最大数，就放一个次大数，针对的是1,1,7这种
		if buf.Len() >= 2 && l1 == arr[2].c && l2 == arr[2].c {
			if arr[1].cnt > 0 {
				p = arr[1]
			}
		} else {
			// 前两次方最大数||穿插了次大数后，再放最大数
			if arr[2].cnt > 0 {
				p = arr[2]
			}
		}
		if p == nil {
			break
		}
		buf.WriteRune(p.c)
		p.cnt--
		l1, l2 = l2, p.c
	}
	return buf.String()
}
