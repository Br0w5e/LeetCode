//179. 最大数
package main


//字符串排序
import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sortNums := make([]string, len(nums))
	for i, num := range nums {
		sortNums[i] = strconv.Itoa(num)
	}
	sort.SliceStable(sortNums, func(i, j int) bool {
		return sortNums[i]+sortNums[j] > sortNums[j]+sortNums[i]
	})

	if sortNums[0] == "0" {
		return "0"
	}
	return strings.Join(sortNums, "")
}