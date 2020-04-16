//合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}	
	sort.Slice(intervals, func (i int, j int) bool  {
		return  intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	for i := 0; i < len(intervals)-1; i++{
		//向后逐步合并
		cur := intervals[i]
		//交叉
		if cur[1] >= intervals[i+1][0] {
			intervals[i+1][0] = cur[0]
			//包含
			if intervals[i+1][1] < cur[1] {
				intervals[i+1][1] = cur[1]
			}
		} else {
			res = append(res, cur)
		}
	}
	res = append(res, intervals[len(intervals)-1])
	return res
}