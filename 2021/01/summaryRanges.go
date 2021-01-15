package main

//228. 汇总区间
import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	pre := nums[0]
	start := pre
	s := strconv.Itoa(pre)
	res := []string{s}
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num != pre+1 {
			if pre != start {
				res[len(res)-1] += "->" + strconv.Itoa(pre)
			}
			start = num
			pre = num
			res = append(res, strconv.Itoa(num))
		} else if i == len(nums)-1 {
			res[len(res)-1] += "->" + strconv.Itoa(num)
		} else {
			pre++
		}
	}
	return res
}