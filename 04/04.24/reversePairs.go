//逆序数组对

//方法一： 暴力法 （超时）
func reversePairs(nums []int) int {
    res := 0
    for i := 0; i < len(nums)-1; i++{
        tmp := 0
        for j := i+1; j < len(nums); j++{
            if nums[i] > nums[j] {
                tmp++
            }
        }
        res += tmp
    }
    return res
}

//归并分治
func reversePairs(nums []int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }
    n := len(nums)
    cnt := 0
    tmp := make([]int, n)
    mergeSort(nums, 0, n-1, &cnt, tmp)
    return cnt
}

func mergeSort(nums []int, left int, right int, cnt *int, tmp []int) {
    //出口
    if left >= right {
        return 
    }
    mid := left+(right-left)/2
    mergeSort(nums, left, mid, cnt, tmp)
    mergeSort(nums, mid+1, right, cnt, tmp)
    merge(nums, left, mid, right, cnt, tmp)
}

func merge(nums []int, left int, mid int, right int, cnt *int, tmp []int) {
    l, r := left, mid+1
    index := left
    for l <= mid && r <= right {
        if nums[l] <= nums[r] {
            tmp[index] = nums[l]
            //左边需要弹出去的时候更新数量
            (*cnt) += r-1-mid
            l++
            index++
        } else {
            tmp[index] = nums[r]
            r++
            index++
        }
    }
    for l <= mid {
        tmp[index] = nums[l]
        (*cnt) += r-1-mid
        l++
        index++
    }
    for r <= right {
        tmp[index] = nums[r]
        r++
        index++
    }
    for i := left; i <= right; i++{
        nums[i] = tmp[i]
    }
}