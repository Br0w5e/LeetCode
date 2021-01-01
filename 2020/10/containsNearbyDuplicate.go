package main
//219. 存在重复元素 II

//暴力记录每个相同元素的下标 到数组中，然后遍历数组
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}
	max := len(nums)
	flag := false
	for _, v := range m {
		if len(v) >= 2 {
			for i := 0; i < len(v)-1; i++ {
				if v[i+1]-v[i] < max {
					max = v[i+1]-v[i]
					flag = true
				}
			}
		}
	}
	return max <= k &&flag
}


//带技巧
func containsNearbyDuplicate1(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		//p为上一个相同元素的位置
		p, ok := m[num]
		if ok && abs(p, i) <= k {
			return true
		}
		//更新value
		m[num] = i
	}
	return false
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

//逆向思考，看范围内有没有满足要求的数字
func containsNearbyDuplicate2(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j:=i+1; j<=i+k && j < len(nums); j++{
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}