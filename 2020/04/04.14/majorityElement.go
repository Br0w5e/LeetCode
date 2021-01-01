//找出数组中 出现次数大于 n/2的数（众数）
func majorityElement(nums []int) int {
    major, count := -1, 0
    for _, num := range nums {
        if count == 0 {
            major = num
        }
        if num == major {
            count += 1
        } else {
            count -= 1
        }
    }
    return major
}

func majorityElement(nums []int) int {
    dic := make(map[int]int, 0)
    n := len(nums) / 2
    for _, v := range nums{
        dic[v]++
    }

    for i, v := range dic {
        if v > n {
            return i
        }
    }
    return 0
}

//找数组中 出现次数大于n/3参考上边的字典思路
func majorityElement(nums []int) []int {
    dic := make(map[int]int, 0)
    n := len(nums) / 3
    res := make([]int, 0)
    
    for _, v := range nums {
        dic[v]++
    }

    for i, v := range dic{
        if v > n {
            res = append(res, i)
        }
    }
    return res
}

func majorityElement(nums []int) []int {
    res := make([]int, 0)
    if nums == nil || len(nums) == 0 {
        return res
    }
    n := len(nums)/3
    cand1 := nums[0]
    cnt1 := 0
    cand2 := nums[0]
    cnt2 := 0

    //摩尔斯投票
    for _, num := range nums {
        if cand1 == num {
            cnt1++
            continue
        }
        if cand2 == num {
            cnt2++ 
            continue 
        }
        if cnt1 == 0 {
            cand1 = num
            cnt1++
            continue
        }
        if cnt2 == 0 {
            cand2 = num
            cnt2++
            continue
        }
        cnt1--
        cnt2--
    }
    //计数阶段
    cnt1 = 0
    cnt2 = 0
    for _, num := range nums {
        switch num {
        case cand1:
            cnt1++
        case cand2:
            cnt2++
        }
    }
    if cnt1 > n {
        res = append(res, cand1)
    }
    if cnt2 > n {
        res = append(res, cand2)
    }
    return res
    
}