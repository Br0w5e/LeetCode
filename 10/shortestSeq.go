package main

//面试题 17.18. 最短超串
func shortestSeq(big []int, small []int) []int {
	min, max, start, end := 0,len(big), 0, 0
	l := len(small)
	smap := make(map[int]int, l)

	for _, v := range small{
		smap[v] = 0
	}
	var meet bool
	for ; end < len(big); end++{
		if _, ok := smap[big[end]]; ok{
			smap[big[end]]++
			if smap[big[end]] > 1{
				continue
			}

			l--
			for l == 0{
				meet = true
				if end - start < max - min{
					max, min = end, start
				}
				_, ok := smap[big[start]]
				if ok {
					smap[big[start]]--
					if smap[big[start]] == 0{
						start++
						l ++
						break
					}
				}
				start++
			}
		}
	}

	if !meet{
		return nil
	}

	return[]int{min, max}
}
