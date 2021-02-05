package main

//424. 替换后的最长重复字符
import "fmt"

//滑动窗口
func characterReplacement(s string, k int) int {
	//记录当前窗口字母出现的个数
	freq := make([]int, 26)
	left, right := 0, 0
	//窗口中数量最多的那一个
	maxCount := 0

	for right < len(s) {
		freq[int(s[right]-'A')]++
		maxCount = max(maxCount, freq[int(s[right]-'A')])
		right++
		// 若当前窗口大小 减去 窗口中最多相同字符的个数 大于 k 时 将窗口最左边的字符 在计数数组中减1
		if right-left > maxCount+k {
			freq[int(s[left]-'A')]--
			left++
		}
	}
	//上边的步骤确定的是窗口的最大差值，当达到差值之后窗口移动就行了
	//虽然这样的操作会导致部分区间不符合条件，即该区间内非最长重复字符超过了k个。但是这样的区间也同样不可能对答案产生贡献。当我们右指针移动到尽头，左右指针对应的区间的长度必然对应一个长度最大的符合条件的区间。
	return right-left
}

func max(a, b int) int   {
	if a > b {
		return a
	}
	return b
}

func main()  {
	s := "AABAABBA"
	k := 1
	fmt.Println(characterReplacement(s, k))
}